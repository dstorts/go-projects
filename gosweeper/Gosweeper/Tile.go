package gosweeper

//this file will define Tile type structs that are extensions of the clickable widgets from gioui
import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Tile struct {
	widget.Clickable
	Content    int //0 for empty, 1-8 for number of adjacent mines, 9 for mine
	IsRevealed bool
	Flagged    bool
}

func (t *Tile) Layout(th *material.Theme, gtx layout.Context) layout.Dimensions {
	// Layout logic for the Tile goes here
	//TODO: func call when this button gets left clicked to reveal tile

	// This is a placeholder implementation
	return layout.Dimensions{}
}
