package main

import (
	"image"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := new(app.Window)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	// Grid size (scalable)
	const gridRows, gridCols = 10, 10
	// Create a 2D slice of clickable widgets
	var cells [gridRows][gridCols]widget.Clickable
	// Track cell states for demonstration
	var cellStates [gridRows][gridCols]string
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.H5(th, "Minesweeper").Layout(gtx)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					cellSize := gtx.Constraints.Max.X / gridCols
					return layout.Grid{
						Axis:      layout.Vertical,
						Columns:   gridCols,
						Alignment: layout.Middle,
					}.Layout(gtx, gridRows*gridCols, func(gtx layout.Context, i int) layout.Dimensions {
						row := i / gridCols
						col := i % gridCols
						var clickedLeft, clickedRight bool
						// Listen for pointer events
						for _, ev := range cells[row][col].Events(gtx) {
							if pev, ok := ev.(pointer.Event); ok {
								if pev.Type == pointer.Press {
									switch pev.Buttons {
									case pointer.ButtonPrimary:
										clickedLeft = true
									case pointer.ButtonSecondary:
										clickedRight = true
									}
								}
							}
						}
						if clickedLeft {
							cellStates[row][col] = "L"
						} else if clickedRight {
							cellStates[row][col] = "R"
						}
						return layout.Stack{}.Layout(gtx,
							layout.Expanded(func(gtx layout.Context) layout.Dimensions {
								return layout.Dimensions{Size: image.Pt(cellSize, cellSize)}
							}),
							layout.Stacked(func(gtx layout.Context) layout.Dimensions {
								// Register for pointer events
								pointer.Rect(image.Rectangle{Max: image.Pt(cellSize, cellSize)}).Add(gtx.Ops)
								pointer.InputOp{
									Tag:   &cells[row][col],
									Types: pointer.Press,
									Grab:  true,
								}.Add(gtx.Ops)
								label := cellStates[row][col]
								return material.Button(th, &cells[row][col], label).Layout(gtx)
							}),
						)
					})
				}),
			)
			e.Frame(gtx.Ops)
		}
	}
}
