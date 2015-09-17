package incoming

import (
	"github.com/Nyarum/noterius/interfaces"
	"github.com/Nyarum/noterius/library/network"
)

type CrumbAuth struct {
}

func (ca *CrumbAuth) PreHandler(netes network.Netes) interfaces.PillDecoder {
	return ca
}

func (ca *CrumbAuth) Process() int {
	return 6
}
