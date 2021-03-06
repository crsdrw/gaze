package algos

import (
	"math/rand"
	"testing"
	"time"

	"github.com/wliao008/gaze"
)

func TestKruskal(t *testing.T) {
	t.Skip("skipping this test for now")
	rand.Seed(time.Now().UTC().UnixNano())
	count := rand.Intn(50) + 1
	for c := 0; c < count; c++ {
		k := NewKruskal(uint16(rand.Intn(100)+1), uint16(rand.Intn(100)+1))
		k.Generate()
		for h := uint16(0); h < k.Board.Height; h++ {
			for w := uint16(0); w < k.Board.Width; w++ {
				if !k.Board.Cells[h][w].IsSet(gaze.VISITED) {
					t.Errorf("Kruskal, every cell should be visited, but not [%d,%d]", h, w)
				}
			}
		}
	}
}

func TestNewKruskal(t *testing.T) {
	k := NewPrim(10, 10)
	for _, row := range k.Board.Cells {
		for _, cell := range row {
			if cell.Flag != 15 {
				t.Errorf("NewKruskal(), every cell should have flag set to 15, got %d", cell.Flag)
			}
		}
	}
}

func BenchmarkKruskalAlgo_1000x500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := NewKruskal(1000, 500)
		k.Generate()
	}
}

func BenchmarkKruskalAlgo_100x50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := NewKruskal(100, 50)
		k.Generate()
	}
}
