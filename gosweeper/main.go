package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand/v2"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// Constants for configuration
const (
	WIDTH     = 800
	HEIGHT    = 600
	GRID_X    = 8 // Columns
	GRID_Y    = 6 // Rows
	NUM_MINES = 5
	TILE_DIM  = 80 // Tile dimension in Dp (pixels)
)

// Defining color constants that will represent different tile states for debugging
// Revealed without a Bomb color will be light gray
var RevealedColor = color.NRGBA{R: 200, G: 200, B: 200, A: 255}

// Unrevealed without a Bomb color will be blue
var UnrevealedColor = color.NRGBA{R: 50, G: 50, B: 200, A: 255}

// Flagged color will be yellow
var FlaggedColor = color.NRGBA{R: 200, G: 200, B: 50, A: 255}

// Revealed with a Bomb color will be red
var RevealedBombColor = color.NRGBA{R: 200, G: 50, B: 50, A: 255}

// Unrevealed with a Bomb color will be purple
var UnrevealedBombColor = color.NRGBA{R: 150, G: 50, B: 150, A: 255}

var BombEmoji = "B" //"\U0001F4A3" // Represents 💣
var FlagEmoji = "⚑" // Flag emoji

// Tile represents a single square in the Minesweeper grid
type Tile struct {
	Clickable     widget.Clickable
	IsRevealed    bool
	IsMine        bool
	AdjacentMines int
}

// Board encapsulates the game grid and logic
type Board struct {
	Rows        int
	Cols        int
	RevealQueue []*Tile
	Tiles       [][]*Tile
}

// NewBoard creates and initializes a new Minesweeper board
func (board *Board) NewBoard(rows, cols, mines int) {
	board.Rows = rows
	board.Cols = cols
	board.RevealQueue = make([]*Tile, 0)
	board.Tiles = make([][]*Tile, rows)
	for y := 0; y < rows; y++ {
		board.Tiles[y] = make([]*Tile, cols)
		for x := 0; x < cols; x++ {
			board.Tiles[y][x] = &Tile{}
		}
	}
	board.GenerateMines(mines)
	board.CalcNeighbors()
}

// GenerateMines randomly assigns a passed number of mines to tiles on the board
func (board *Board) GenerateMines(mineCount int) {
	// Create a flat list of booleans representing mine placement
	// We'll take advantage of dynamic array sizing by appending new bools
	var mineList []bool
	for i := 0; i < board.Rows*board.Cols; i++ {
		if i < mineCount {
			mineList = append(mineList, true)
		} else {
			mineList = append(mineList, false)
		}
	}
	// By keeping the mine placement in a flat array, we can shuffle it easily and then map it back to the 2D grid
	rand.Shuffle(len(mineList), func(i, j int) {
		mineList[i], mineList[j] = mineList[j], mineList[i]
	})
	// Map the shuffled mine list back to the 2D grid
	for i := 0; i < len(mineList); i++ {
		row_index := int(i % board.Cols)
		col_index := int(i / board.Cols)
		board.Tiles[col_index][row_index].IsMine = mineList[i]
		//fmt.Printf("Tile[%d][%d] Address: %p\n", col_index, row_index, &board.Tiles[col_index][row_index])
	}
}

// CalcNeighbors will look at all 8 adjacent tiles from any given tile and set its AdjacentMines field
func (board *Board) CalcNeighbors() {
	for row := 0; row < board.Rows; row++ {
		for col := 0; col < board.Cols; col++ {
			isTopEdge := row == 0
			isBottomEdge := row == (board.Rows - 1)
			isLeftEdge := col == 0
			isRightEdge := col == (board.Cols - 1)

			// Thx Jake.
			if !isLeftEdge && !isTopEdge {
				// here if above-left neighbor exists
				if board.Tiles[row-1][col-1].IsMine {
					board.Tiles[row][col].AdjacentMines++
				}
			}
			if !isTopEdge {
				// here if directly-above neighbor exists
				if board.Tiles[row-1][col].IsMine {
					board.Tiles[row][col].AdjacentMines++
				}
			}
			if !isTopEdge && !isRightEdge {
				// here if above-right neighbor exists
				if board.Tiles[row-1][col+1].IsMine {
					board.Tiles[row][col].AdjacentMines++
				}
			}
			if !isLeftEdge {
				// here if directly-left neighbor exists
				if board.Tiles[row][col-1].IsMine {
					board.Tiles[row][col].AdjacentMines++
				}
			}
			if !isRightEdge {
				// here if directly-right neighbor exists
				if board.Tiles[row][col+1].IsMine {
					board.Tiles[row][col].AdjacentMines++
				}
			}
			if !isLeftEdge && !isBottomEdge {
				// here if below-left neighbor exists
				if board.Tiles[row+1][col-1].IsMine {
					board.Tiles[row][col].AdjacentMines++
				}
			}
			if !isBottomEdge {
				// here if directly-below neighbor exists
				if board.Tiles[row+1][col].IsMine {
					board.Tiles[row][col].AdjacentMines++
				}
			}
			if !isBottomEdge && !isRightEdge {
				// here if below-right neighbor exists
				if board.Tiles[row+1][col+1].IsMine {
					board.Tiles[row][col].AdjacentMines++
				}
			}
		}
	}
}

func (board *Board) Reveal(row, col int) {
	// do nothing if out of index ranges
	if row < 0 || row >= board.Rows || col < 0 || col >= board.Cols {
		fmt.Printf("Tile[%d][%d] was OB\n", row, col)
		return
	}
	currTile := board.Tiles[row][col]
	if currTile.IsRevealed || currTile.IsMine {
		fmt.Printf("Tile[%d][%d] was already revealed\n", row+1, col+1)
		return
	}

	// mark tile as revealed
	currTile.IsRevealed = true
	fmt.Printf("Recursively Revealed Tile[%d][%d]\n", row+1, col+1)

	// if current tile's bomb count is zero, reveal all its neighbors
	if currTile.AdjacentMines == 0 {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dy == 0 && dx == 0 {
					continue // skip this loop iteration if we are trying to index the tile itself
				}
				board.Reveal(row+dy, col+dx)
			}
		}
	}
}

func (board *Board) RevealV2(row, col int) {
	//recursively check every adjacent tile for
	board.BuildRevealQueue(row, col)
	fmt.Printf("Len of RevealQueue: %d\n", len(board.RevealQueue))
	for ti := 0; ti < len(board.RevealQueue); ti++ {
		board.RevealQueue[ti].IsRevealed = true
	}
	board.RevealQueue = make([]*Tile, 0)
}

func (board *Board) BuildRevealQueue(row, col int) {
	//catch out of bounds indicies on every call instead of trying to avoid them
	if row < 0 || row >= board.Rows || col < 0 || col >= board.Cols {
		return
	}
	//stop working with this current tile if it was already revealed or is a bomb
	if board.Tiles[row][col].IsRevealed || board.Tiles[row][col].IsMine {
		return
	}
	for i := 0; i < len(board.RevealQueue); i++ {
		if board.Tiles[row][col] == board.RevealQueue[i] {
			//here if current tile is alredy in the reveal queue
			//don't add it twice
			return
		}
	}
	board.RevealQueue = append(board.RevealQueue, board.Tiles[row][col])
	if board.Tiles[row][col].AdjacentMines == 0 {
		for delta_y := -1; delta_y <= 1; delta_y++ {
			for delta_x := -1; delta_x <= 1; delta_x++ {
				if delta_y == 0 && delta_x == 0 {
					//skip the self
					//otherwise infinite recursion
					continue
				}
				board.BuildRevealQueue(row+delta_y, col+delta_x)
			}
		}
	}
}

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("GOsweeper (a Minesweeper Game...)"), app.Size(unit.Dp(WIDTH), unit.Dp(HEIGHT)))

		// Initialize the board using our new struct
		var gameBoard Board
		gameBoard.NewBoard(GRID_Y, GRID_X, NUM_MINES)

		if err := run(w, &gameBoard); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window, board *Board) error {
	th := material.NewTheme()
	var ops op.Ops

	for {
		e := w.Event()
		switch e := e.(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				var rows []layout.FlexChild
				for y := 0; y < board.Rows; y++ {
					// Capture current row to pass to renderer
					rows = append(rows, renderRow(gtx, th, board, y))
				}

				return layout.Flex{Axis: layout.Vertical}.Layout(gtx, rows...)
			})

			e.Frame(gtx.Ops)
		}
	}
}

func renderRow(gtx layout.Context, th *material.Theme, board *Board, row int) layout.FlexChild {
	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		var children []layout.FlexChild
		for x := 0; x < board.Cols; x++ {
			tile := board.Tiles[row][x]
			children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if tile.Clickable.Clicked(gtx) {
					if tile.IsMine {
						tile.IsRevealed = true
					} else {
						board.RevealV2(row, x)
					}
				}

				btn := material.Button(th, &tile.Clickable, "")
				if tile.IsRevealed {
					if tile.IsMine {
						btn.Text = BombEmoji
						btn.Background = RevealedBombColor
					} else {
						btn.Text = fmt.Sprintf("%d", tile.AdjacentMines)
						btn.Background = RevealedColor
					}

				} else {
					if tile.IsMine {
						btn.Text = BombEmoji
						btn.Background = UnrevealedBombColor
					} else {
						btn.Text = fmt.Sprintf("%d", tile.AdjacentMines)
						btn.Background = UnrevealedColor
					}
				}

				gtx.Constraints.Min.X = gtx.Dp(TILE_DIM)
				gtx.Constraints.Min.Y = gtx.Dp(TILE_DIM)
				return layout.UniformInset(unit.Dp(1)).Layout(gtx, btn.Layout)
			}))
		}
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, children...)
	})
}
