package userstore

import (
	"exercise/src/types/common"
)

type canvasDrawingImpl struct {
	canvas    common.Canvas
	maxHeight int
	maxWidth  int
}

func NewCanvasDrawing(height, width int) (*canvasDrawingImpl, error) {
	cd := &canvasDrawingImpl{
		canvas:    make(common.Canvas, height),
		maxHeight: height,
		maxWidth:  width,
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
		for x := 0; x < cd.maxWidth; x++ {
			cd.canvas[y][x] = ' '
		}
	}
	return nil
}

func (cd *canvasDrawingImpl) FillCanvas(rect common.Rectangle) error {
	if err := cd.isValidRectangle(rect); err != nil {
		return err
	}

	if rect.Outline == nil && rect.Fill == nil {
		return &EmptyDrawingParamsError{}
	} else if rect.Outline == nil {
		cd.fill(rect)
	} else if rect.Fill == nil {
		cd.fillOutline(rect)
	} else if *rect.Outline == *rect.Fill {
		cd.fill(rect)
	} else {
		cd.fillWithOutline(rect)
	}
	return nil
}

// isValidRectangle checks if a rectangle drawing is valid and
// returns an error if not
func (cd *canvasDrawingImpl) isValidRectangle(rect common.Rectangle) error {
	if rect.Location.X < 0 || rect.Location.Y < 0 ||
		rect.Location.X > cd.maxWidth || rect.Location.Y > cd.maxHeight {
		return &RectangleOutOfBoundsError{}
	}

	if rect.Height < 0 || rect.Location.Y+rect.Height > cd.maxHeight ||
		rect.Width < 0 || rect.Location.X+rect.Width > cd.maxWidth {
		return &DrawingOutOfBoundsError{}
	}

	return nil
}

// fill draws the entire rectangle
func (cd *canvasDrawingImpl) fill(rect common.Rectangle) {
	upperMostY := rect.Location.Y
	lowerMostY := rect.Location.Y + rect.Height
	leftMostX := rect.Location.X
	rightMostX := rect.Location.X + rect.Width

	for y := upperMostY; y < lowerMostY; y++ {
		for x := leftMostX; x < rightMostX; x++ {
			cd.canvas[y][x] = *rect.Fill
		}
	}
}

// fillOutline draws only the ouline of a rectangle
func (cd *canvasDrawingImpl) fillOutline(rect common.Rectangle) {
	upperMostY := rect.Location.Y
	lowerMostY := rect.Location.Y + rect.Height
	leftMostX := rect.Location.X
	rightMostX := rect.Location.X + rect.Width

	for y := upperMostY; y < lowerMostY; y++ {
		cd.canvas[y][leftMostX] = *rect.Outline
		cd.canvas[y][rightMostX-1] = *rect.Outline
		// if this position is not a border of the rectangle, skip it
		if y != upperMostY && y != lowerMostY-1 {
			continue
		}
		for x := leftMostX; x < rightMostX; x++ {
			cd.canvas[y][x] = *rect.Outline
		}
	}
}

// fillWithOutline draws the rectangle and then draws its outline
func (cd *canvasDrawingImpl) fillWithOutline(rect common.Rectangle) {
	cd.fill(rect)
	cd.fillOutline(rect)
}

func (cd *canvasDrawingImpl) GetCanvas() common.Canvas {
	return cd.canvas
}
