package main

//DrawTheCell
//1. draw the leaves
//2. combine the child if they are from same node and draw the line

func DrawTheCell(vesicle *Vesicle) {
	vesicleDepth := MeasureDepth(vesicle)
	if vesicle != nil {
		for _, childVesicle := range vesicle.vesicles {
			if childVesicle.vesicles == nil {
				DrawLeaves(vesicle)
			} else {
				DrawVesicle(vesicle)

			}
		}
	}
}

func DrawLeaves(vesicle *Vesicle) {

}

func MeasureDepth(vesicle *Vesicle) int {

	return 0
}

func DrawVesicle(vesicle *Vesicle) {

}

func CombineCanvas() {

}

func CreateCanvas() {

}

/* This is the drawing code from spatial

func DrawGameBoard(board GameBoard, cellWidth int) Canvas {
	height := len(board) * cellWidth
	width := len(board[0]) * cellWidth
	c := CreateNewCanvas(width, height)

	// declare colors
	blue := MakeColor(0, 0, 255)
	red := MakeColor(255, 0, 0)


		// draw the grid lines in white
		c.SetStrokeColor(white)
		DrawGridLines(c, cellWidth)


	// fill in colored squares
	for i := range board {
		for j := range board[i] {
			if board[i][j].strategy == "C" {
				c.SetFillColor(blue)
			} else if board[i][j].strategy == "D" {
				c.SetFillColor(red)
			} else {
				panic("Error: Out of range value " + string(board[i][j].strategy) + " in board when drawing board.")
			}
			x := j * cellWidth
			y := i * cellWidth
			c.ClearRect(x, y, x+cellWidth, y+cellWidth)
			c.Fill()
		}
	}
	return c
}

func DrawGameBoardGif(board GameBoard, cellWidth int) image.Image {
	height := len(board) * cellWidth
	width := len(board[0]) * cellWidth
	c := CreateNewCanvas(width, height)

	blue := MakeColor(0, 0, 255)
	red := MakeColor(255, 0, 0)
	for i := range board {
		for j := range board[i] {
			if board[i][j].strategy == "C" {
				c.SetFillColor(blue)
			} else if board[i][j].strategy == "D" {
				c.SetFillColor(red)
			} else {
				panic("Error: Out of range value " + string(board[i][j].strategy) + " in board when drawing board.")
			}
			x := j * cellWidth
			y := i * cellWidth
			c.ClearRect(x, y, x+cellWidth, y+cellWidth)
			c.Fill()
		}
	}
	return c.img
}

*/
