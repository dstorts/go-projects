package gosweeper

//this file will define the grid that holds the tiles for minesweeper where the grid will extend a gioui container window

import (
	"gioui.org/app"
)

type Minesweeper_Grid struct {
	app.Window
	Rows    uint8
	Columns uint8
	Tiles   [][]Tile
}

//now I will define an interface for the
