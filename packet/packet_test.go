package packet

import (
	"testing"

	"github.com/Nyarum/noterius/core"
	"github.com/Nyarum/noterius/database"
	"github.com/Nyarum/noterius/entitie"
	"github.com/Nyarum/noterius/manager"
)

func BenchmarkEncode940Packet(b *testing.B) {
	buffers := core.NewBuffers()
	player := entitie.NewPlayer(buffers)
	database := database.NewDatabase(&core.DatabaseInfo{})
	manager := manager.NewManager(database)
	packet := NewPacket(player, manager)

	for n := 0; n < b.N; n++ {
		packet.Encode(940)
	}
}

func BenchmarkEncode931Packet(b *testing.B) {
	buffers := core.NewBuffers()
	player := entitie.NewPlayer(buffers)
	database := database.NewDatabase(&core.DatabaseInfo{})
	manager := manager.NewManager(database)
	packet := NewPacket(player, manager)

	for n := 0; n < b.N; n++ {
		packet.Encode(931)
	}
}
