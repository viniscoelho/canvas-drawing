package userstore

import (
	"exercise/src/types/common"
)

type canvasDrawingImpl struct {
	canvas    common.Canvas
	maxWidth  int
	maxHeight int
}

func NewCanvasDrawing(width, height int) (*canvasDrawingImpl, error) {
	cd := &canvasDrawingImpl{
		canvas:    make(common.Canvas, height),
		maxWidth:  width,
		maxHeight: height,
	}

	err := cd.initializeCanvas()
	if err != nil {
		return nil, err
	}

	return cd, nil
}

func (cd *canvasDrawingImpl) initializeCanvas() error {
	for y := 0; y < cd.maxHeight; y++ {
		cd.canvas[y] = make([]rune, cd.maxWidth)
		for x := 0; x < cd.maxHeight; x++ {
			cd.canvas[y][x] = ' '
		}
	}
	return nil
}

func (cd *canvasDrawingImpl) FillCanvas(pos common.Rectangle, fill, outline rune) error {
	if pos.Location.X < 0 || pos.Location.Y < 0 ||
		pos.Location.X > cd.maxWidth || pos.Location.Y > cd.maxHeight {
		return &RectangleOutOfBoundsError{}
	}

	if pos.Height < 0 || pos.Location.Y+pos.Height > cd.maxHeight ||
		pos.Width < 0 || pos.Location.X+pos.Width > cd.maxWidth {
		return &DrawingOutOfBoundsError{}
	}

	for y := pos.Location.Y; y < pos.Location.Y+pos.Height; y++ {
		for x := pos.Location.X; x < pos.Location.X+pos.Width; x++ {
			cd.canvas[y][x] = fill
		}
	}
	return nil
}

func (cd *canvasDrawingImpl) GetCanvas() common.Canvas {
	return cd.canvas
}
