package entities

import (
	"database/sql"
	"strings"

	"github.com/Nyarum/noterius/network/common"
	"github.com/Nyarum/noterius/network/in"

	"go.uber.org/zap"

	kallax "gopkg.in/src-d/go-kallax.v1"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/noterius/common/crypt"
	"github.com/Nyarum/noterius/models"
	"github.com/Nyarum/noterius/network/out"
)

type Player struct {
	DB           *sql.DB
	World        *actor.PID
	PacketSender *actor.PID
	Logger       *zap.SugaredLogger

	Info *models.Player
	Time string
}

func (state *Player) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Stopping:
		if state.Info != nil {
			state.Info.IsActive = false
			models.NewPlayerStore(state.DB).Update(state.Info)
		}
	case RecordTime:
		state.Time = msg.Time

		state.PacketSender.Tell(SendPacket{
			Packet: &out.Date{
				Time: msg.Time,
			},
		})
	case *in.Auth:
		var (
			playerStore = models.NewPlayerStore(state.DB)
		)

		getPlayer, err := playerStore.FindOne(
			models.NewPlayerQuery().FindByUsername(msg.Login).WithCharacters(nil),
		)
		if err != nil {
			if err == kallax.ErrNotFound {
				state.PacketSender.Tell(SendPacketWithLogout{
					Packet: &out.Auth{
						ErrorCode: common.PlayerIsNotFound.GetID(),
					},
				})
			}

			state.Logger.Errorw("Find error", "error", err)
			return
		}

		encryptPassword, err := crypt.EncryptPassword(strings.ToUpper(getPlayer.Password[:24]), state.Time)
		if err != nil {
			state.Logger.Errorw("Encrypt password error", "error", err)
			return
		}

		if encryptPassword != msg.Password {
			state.Logger.Debugw("Verify error", "username", msg.Login, "error", common.PasswordIncorrect)
			state.PacketSender.Tell(SendPacketWithLogout{
				Packet: &out.Auth{
					ErrorCode: common.PasswordIncorrect.GetID(),
				},
			})
			return
		}

		if getPlayer.IsActive {
			state.Logger.Debugw("Account is already logged", "username", msg.Login)
			state.PacketSender.Tell(SendPacketWithLogout{
				Packet: &out.Auth{
					ErrorCode: common.PlayerInGame.GetID(),
				},
			})
			return
		}

		authPacket := &out.Auth{}
		authPacket.SetPincode(getPlayer.Pincode)
		for _, character := range getPlayer.Characters {
			charSub := common.CharacterSub{
				Name:  character.Name,
				Job:   character.Job,
				Level: character.Level,
				Look: common.CharacterLookSub{
					Race: character.Race,
				},
			}
			charSub.SetFlag(character.Enabled)

			authPacket.Characters = append(authPacket.Characters, charSub)
		}

		state.Info = getPlayer

		state.PacketSender.Tell(SendPacket{
			Packet: authPacket,
		})

		getPlayer.IsActive = true
		playerStore.Update(state.Info)
	case *in.NewSecret:
		state.Info.Pincode = &msg.Password

		playerStore := models.NewPlayerStore(state.DB)
		playerStore.Update(state.Info)

		state.PacketSender.Tell(SendPacket{
			Packet: &out.NewSecret{
				ErrorCode: 0,
			},
		})
	case *in.ChangeSecret:
		resp := &out.ChangeSecret{}

		if *state.Info.Pincode != msg.PasswordOld {
			state.Logger.Debugw("Old secret password is incorrect", "username", state.Info.Username)

			resp.ErrorCode = common.SecretPasswordIncorrect.GetID()
			state.PacketSender.Tell(SendPacket{
				Packet: resp,
			})
			return
		}

		playerStore := models.NewPlayerStore(state.DB)
		playerStore.Update(state.Info)

		state.PacketSender.Tell(SendPacket{
			Packet: resp,
		})
	case *in.DeleteCharacter:
		resp := &out.DeleteCharacter{}

		if *state.Info.Pincode != msg.Secret {
			state.Logger.Debugw("Secret password is incorrect", "username", state.Info.Username)

			resp.ErrorCode = common.SecretPasswordIncorrect.GetID()
			state.PacketSender.Tell(SendPacket{
				Packet: resp,
			})
			return
		}

		for n, char := range state.Info.Characters {
			if char.Name == msg.Name {
				state.Info.Characters[n].Enabled = false

				characterStore := models.NewCharacterStore(state.DB)
				characterStore.Update(state.Info.Characters[n])
			}
		}

		state.PacketSender.Tell(SendPacket{
			Packet: resp,
		})
	case *in.CreateCharacter:
		var (
			resp           = &out.CreateCharacter{}
			characterStore = models.NewCharacterStore(state.DB)
			mapStore       = models.NewMapStore(state.DB)
		)

		charGet, err := characterStore.FindOne(
			models.NewCharacterQuery().FindByName(msg.Name),
		)
		if charGet != nil {
			resp.ErrorCode = common.ExistCharName.GetID()
			state.PacketSender.Tell(SendPacket{
				Packet: resp,
			})
			return
		}

		mapGet, err := mapStore.FindOne(
			models.NewMapQuery().FindByName(msg.Map),
		)
		if err != nil {
			if err == kallax.ErrNotFound {
				resp.ErrorCode = common.InvalidBirthLocation.GetID()
			} else {
				resp.ErrorCode = common.InternalError.GetID()
			}

			state.PacketSender.Tell(SendPacket{
				Packet: resp,
			})

			state.Logger.Errorw("Find error", "error", err)
			return
		}

		characterModel := models.NewCharacter()
		characterModel.Player = state.Info
		characterModel.Name = msg.Name
		characterModel.Map = mapGet
		characterModel.Job = "Newbie"
		characterModel.Race = msg.Look.Race
		characterModel.Level = 1
		characterModel.Enabled = true

		err = characterStore.Insert(characterModel)
		if err != nil {
			resp.ErrorCode = common.InternalError.GetID()
			state.PacketSender.Tell(SendPacket{
				Packet: resp,
			})

			state.Logger.Errorw("Insert error", "error", err)
			return
		}

		state.Info.Characters = append(state.Info.Characters, characterModel)

		state.PacketSender.Tell(SendPacket{
			Packet: &out.CreateCharacter{},
		})
	}
}
