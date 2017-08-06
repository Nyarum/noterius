package entities

import (
	"database/sql"
	"strings"

	"go.uber.org/zap"

	kallax "gopkg.in/src-d/go-kallax.v1"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/noterius/common/crypt"
	"github.com/Nyarum/noterius/models"
	"github.com/Nyarum/noterius/network/errors"
	"github.com/Nyarum/noterius/network/out"
)

type Player struct {
	DB           *sql.DB
	World        *actor.PID
	PacketSender *actor.PID
	Logger       *zap.SugaredLogger

	Time string
}

func (state *Player) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case RecordTime:
		state.Time = msg.Time

		state.PacketSender.Tell(SendPacket{
			Packet: &out.Date{
				Time: msg.Time,
			},
		})
	case Auth:
		getPlayer, err := models.NewPlayerStore(state.DB).FindOne(
			models.NewPlayerQuery().FindByUsername(msg.Login),
		)
		if err != nil {
			if err == kallax.ErrNotFound {
				state.PacketSender.Tell(SendPacketWithLogout{
					Packet: &out.Auth{
						ErrorCode: errors.PlayerIsNotFound.GetID(),
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
			state.Logger.Debugw("Verify error", "username", msg.Login, "error", errors.PasswordIncorrect)
			state.PacketSender.Tell(SendPacketWithLogout{
				Packet: &out.Auth{
					ErrorCode: errors.PasswordIncorrect.GetID(),
				},
			})
			return
		}

		state.PacketSender.Tell(SendPacket{
			Packet: (&out.Auth{}).SetTestData(),
		})
	case Logout:
		// Something we do with database and other services
		// and exit

		state.PacketSender.Tell(msg)
	}
}
