package incoming

import (
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/interface"
	"github.com/Nyarum/noterius/library/network"

	"log"
)

type AuthCrumb struct {
	Key           string
	Login         string
	Password      string
	MAC           string
	IsCheat       uint16
	ClientVersion uint16
}

func (a *AuthCrumb) PreHandler(netes network.Netes) interfaces.PillDecoder {
	netes.ReadString(&a.Key)
	netes.ReadString(&a.Login)
	netes.ReadString(&a.Password)
	netes.ReadString(&a.MAC)
	netes.ReadUint16(&a.IsCheat)
	netes.ReadUint16(&a.ClientVersion)

	return a
}

func (a *AuthCrumb) Process(player entitie.Player) ([]int, error) {
	log.Println("Login:", a.Login)
	log.Println("Client version:", a.ClientVersion)

	return []int{931}, nil
}
