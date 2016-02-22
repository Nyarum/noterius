package packet

import (
	"errors"
	"time"

	"github.com/Nyarum/noterius/database"
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/library/network"
	"github.com/Nyarum/noterius/support"
	log "github.com/Sirupsen/logrus"
)

type PacketFactory interface {
	Packet(*database.Database) (func(network.Netes), func(*entitie.Player) []int)
}

type Packet struct {
	pills    map[int]PacketFactory
	Player   *entitie.Player
	Database *database.Database
}

type PacketHeader struct {
	Len      uint16
	UniqueId uint32
	Opcode   uint16
}

func NewPacket(player *entitie.Player, database *database.Database) *Packet {
	return &Packet{
		pills: map[int]PacketFactory{
			support.OP_CLIENT_LOGIN:            &IncomingAuth{},
			support.OP_CLIENT_LOGOUT:           &IncomingExit{},
			support.OP_SERVER_LOGIN:            &OutcomingAuth{},
			support.OP_SERVER_CHAPSTR:          &OutcomingDate{},
			support.OP_CLIENT_DELCHA:           &IncomingDeleteCharacter{},
			support.OP_SERVER_DELCHA:           &OutcomingDeleteCharacter{},
			support.OP_CLIENT_UPDATE_PASSWORD2: &IncomingUpdatePassword{},
			support.OP_SERVER_UPDATE_PASSWORD2: &OutcomingUpdatePassword{},
			support.OP_CLIENT_NEWCHA:           &IncomingNewCharacter{},
			support.OP_SERVER_NEWCHA:           &OutcomingNewCharacter{},
			support.OP_CLIENT_BGNPLAY:          &IncomingBeginPlay{},
			support.OP_SERVER_ENTERMAP:         &OutcomingEnterMap{},
			support.OP_CLIENT_KITBAGTEMP_SYNC:  &IncomingKitbagtempSync{},
		},
		Player:   player,
		Database: database,
	}
}

func (p *Packet) GetPck(opcode int) (PacketFactory, error) {
	get, ok := p.pills[opcode]
	if !ok {
		return nil, errors.New("Packet is not found")
	}

	return get, nil
}

func (p *Packet) Encode(opcode int) ([]byte, error) {
	start := time.Now()
	netes := network.NewParser([]byte{})

	pck, err := p.GetPck(opcode)
	if err != nil {
		return nil, err
	}

	handler, process := pck.Packet(p.Database)

	process(p.Player)

	handler(netes)
	err = netes.Error()
	if err != nil {
		return nil, err
	}

	data := string(netes.Bytes())
	netes.Reset()

	header := PacketHeader{Len: uint16(len(data) + 8), UniqueId: 128, Opcode: uint16(opcode)}

	netes.SetEndian(network.LittleEndian).WriteUint16(header.Len)
	netes.SetEndian(network.BigEndian).WriteUint32(header.UniqueId)
	netes.SetEndian(network.LittleEndian).WriteUint16(header.Opcode)
	netes.WriteBytes([]byte(data))

	log.Info("Elapled time in Encode function: ", time.Since(start))

	return netes.Bytes(), nil
}

func (p *Packet) Decode(buf []byte) ([]int, error) {
	var (
		header PacketHeader    = PacketHeader{}
		netes  *network.Parser = network.NewParser(buf)
	)

	netes.SetEndian(network.LittleEndian).ReadUint16(&header.Len)
	netes.SetEndian(network.BigEndian).ReadUint32(&header.UniqueId)
	netes.SetEndian(network.LittleEndian).ReadUint16(&header.Opcode)

	pck, err := p.GetPck(int(header.Opcode))
	if err != nil {
		return nil, err
	}

	handler, process := pck.Packet(p.Database)

	handler(netes)
	err = netes.Error()
	if err != nil {
		return nil, err
	}

	opcodes := process(p.Player)

	return opcodes, nil
}
