package pill

import (
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/library/network"
	log "github.com/Sirupsen/logrus"
)

type Auth struct {
	Key           string
	Login         string
	Password      string
	MAC           string
	IsCheat       uint16
	ClientVersion uint16
	err           error
	opcodes       []int
}

func (p *Auth) Error() error {
	return p.err
}

func (p *Auth) Opcodes() []int {
	return p.opcodes
}

func (p *Auth) Handler(netes network.Netes) PillFactory {
	netes.ReadString(&p.Key)
	netes.ReadString(&p.Login)
	netes.ReadString(&p.Password)
	netes.ReadString(&p.MAC)
	netes.ReadUint16(&p.IsCheat)
	netes.ReadUint16(&p.ClientVersion)

	return p
}

func (p *Auth) Process(player *entitie.Player) PillFactory {
	log.WithFields(log.Fields{
		"value": p.Login,
	}).Info("Login")
	log.WithFields(log.Fields{
		"value": p.ClientVersion,
	}).Info("Client version")

	p.opcodes = []int{CharacterListOpcode}

	return p
}
