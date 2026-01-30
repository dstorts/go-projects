package main

import (
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	WIDTH  = 10
	HEIGHT = 10
	BOMBS  = 15
)

// Tile extends a button to track game state
type Tile struct {
	*widget.Button
	IsBomb        bool
	NeighborCount int
	Revealed      bool
	X, Y          int
}

type Board struct {
	Width, Height int
	Tiles         [][]*Tile
}

func NewBoard(w, h int) *Board {
	b := &Board{Width: w, Height: h}
	b.Tiles = make([][]*Tile, h)
	for y := 0; y < h; y++ {
		b.Tiles[y] = make([]*Tile, w)
	}
	return b
}

func (b *Board) GenerateTiles(onTap func(*Tile)) {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			t := &Tile{X: x, Y: y}
			// Use a closure to capture the specific tile
			t.Button = widget.NewButton("", func() { onTap(t) })
			b.Tiles[y][x] = t
		}
	}
}

func (b *Board) GenerateBombs(count int) {
	rand.Seed(time.Now().UnixNano())
	placed := 0
	for placed < count {
		x, y := rand.Intn(b.Width), rand.Intn(b.Height)
		if !b.Tiles[y][x].IsBomb {
			b.Tiles[y][x].IsBomb = true
			placed++
		}
	}
}

func (b *Board) CalcNeighbors() {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			if b.Tiles[y][x].IsBomb {
				continue
			}
			count := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					ny, nx := y+dy, x+dx
					if ny >= 0 && ny < b.Height && nx >= 0 && nx < b.Width && b.Tiles[ny][nx].IsBomb {
						count++
					}
				}
			}
			b.Tiles[y][x].NeighborCount = count
		}
	}
}

func revealTile(t *Tile) {
	if t.Revealed {
		return
	}
	t.Revealed = true
	if t.IsBomb {
		t.Button.SetText("B")
		// In Fyne, colors are often managed via themes.
		// For a specific button, you can swap it with a colored object.
	} else {
		if t.NeighborCount > 0 {
			t.Button.SetText(strconv.Itoa(t.NeighborCount))
		}
	}
	t.Button.Disable() // Grey out and prevent further clicks
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Minesweeper Go")

	board := NewBoard(WIDTH, HEIGHT)
	grid := container.New(layout.NewGridLayout(WIDTH))

	board.GenerateTiles(func(t *Tile) {
		revealTile(t)
	})
	board.GenerateBombs(BOMBS)
	board.CalcNeighbors()

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			grid.Add(board.Tiles[y][x])
		}
	}

	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(500, 400))
	myWindow.ShowAndRun()
}
