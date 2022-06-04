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
	for r := 0; r < cd.maxHeight; r++ {
		cd.canvas[r] = make([]string, cd.maxWidth)
	}
	return nil
}

func (cd *canvasDrawingImpl) FillCanvas(position common.Rectangle, fill, outline rune) error {
	for r := 0; r < cd.maxHeight; r++ {
		cd.canvas[r] = make([]string, cd.maxWidth)
	}
	return nil
}

func (cd *canvasDrawingImpl) GetCanvas() common.Canvas {
	return cd.canvas
}
