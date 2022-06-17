package canvas

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"canvas-drawing/src/types/common"
)

func TestCanvasDrawing_Fill(t *testing.T) {
	assert := assert.New(t)

	cd, err := NewCanvasDrawing(15, 30)
	assert.NoError(err)

	outline := '@'
	fill := 'X'
	rectA := common.Rectangle{
		Location: common.Coordinates{
			X: 3,
			Y: 2,
		},
		Height:  3,
		Width:   5,
		Outline: &outline,
		Fill:    &fill,
	}
	err = cd.FillCanvas(rectA)
	assert.NoError(err)

	outline = 'X'
	fill = 'O'
	rectB := common.Rectangle{
		Location: common.Coordinates{
			X: 10,
			Y: 3,
		},
		Height:  6,
		Width:   14,
		Outline: &outline,
		Fill:    &fill,
	}
	err = cd.FillCanvas(rectB)
	assert.NoError(err)

	expectedCanvas := []string{
		"",
		"",
		"   @@@@@",
		"   @XXX@  XXXXXXXXXXXXXX",
		"   @@@@@  XOOOOOOOOOOOOX",
		"          XOOOOOOOOOOOOX",
		"          XOOOOOOOOOOOOX",
		"          XOOOOOOOOOOOOX",
		"          XXXXXXXXXXXXXX",
	}

	canvas := cd.GetCanvas()
	for y := 0; y < len(expectedCanvas); y++ {
		// remove canvas trailing spaces to facilitate comparison with test case
		canvasRow := strings.TrimRight(canvas[y], " ")
		assert.Equal(expectedCanvas[y], canvasRow, "canvas does not match")
	}
}

func TestCanvasDrawing_OverlapingFill(t *testing.T) {
	assert := assert.New(t)

	cd, err := NewCanvasDrawing(15, 30)
	assert.NoError(err)

	fill := '.'
	rectA := common.Rectangle{
		Location: common.Coordinates{
			X: 14,
			Y: 0,
		},
		Height: 6,
		Width:  7,
		Fill:   &fill,
	}
	err = cd.FillCanvas(rectA)
	assert.NoError(err)

	outline := 'O'
	rectB := common.Rectangle{
		Location: common.Coordinates{
			X: 0,
			Y: 3,
		},
		Height:  4,
		Width:   8,
		Outline: &outline,
	}
	err = cd.FillCanvas(rectB)
	assert.NoError(err)

	outline = 'X'
	fill = 'X'
	rectC := common.Rectangle{
		Location: common.Coordinates{
			X: 5,
			Y: 5,
		},
		Height:  3,
		Width:   5,
		Outline: &outline,
		Fill:    &fill,
	}
	err = cd.FillCanvas(rectC)
	assert.NoError(err)

	expectedCanvas := []string{
		"              .......",
		"              .......",
		"              .......",
		"OOOOOOOO      .......",
		"O      O      .......",
		"O    XXXXX    .......",
		"OOOOOXXXXX",
		"     XXXXX",
	}

	canvas := cd.GetCanvas()
	for y := 0; y < len(expectedCanvas); y++ {
		// remove canvas trailing spaces to facilitate comparison with test case
		canvasRow := strings.TrimRight(canvas[y], " ")
		assert.Equal(expectedCanvas[y], canvasRow, "canvas does not match")
	}
}

func TestCanvasDrawing_SinglePoint(t *testing.T) {
	assert := assert.New(t)

	cd, err := NewCanvasDrawing(15, 30)
	assert.NoError(err)

	fill := '.'
	rectA := common.Rectangle{
		Location: common.Coordinates{
			X: 0,
			Y: 0,
		},
		Height: 1,
		Width:  1,
		Fill:   &fill,
	}
	err = cd.FillCanvas(rectA)
	assert.NoError(err)

	expectedCanvas := []string{
		".",
		"",
		"",
		"",
	}

	canvas := cd.GetCanvas()
	for y := 0; y < len(expectedCanvas); y++ {
		// remove canvas trailing spaces to facilitate comparison with test case
		canvasRow := strings.TrimRight(canvas[y], " ")
		assert.Equal(expectedCanvas[y], canvasRow, "canvas does not match")
	}
}

func TestCanvasDrawing_InvalidCanvas(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCanvasDrawing(0, 0)
	assert.Error(err)
}

func TestCanvasDrawing_InvalidRectangles(t *testing.T) {
	assert := assert.New(t)

	cd, err := NewCanvasDrawing(10, 10)
	assert.NoError(err)

	rectangles := generateInvalidRectangles()
	for _, rect := range rectangles {
		err := cd.FillCanvas(rect)
		assert.Error(err)
	}
}

func generateInvalidRectangles() []common.Rectangle {
	rectangles := make([]common.Rectangle, 0)

	// negative X-index
	rect := common.NewDefaultRectangle()
	rect.Location.X = -1
	rectangles = append(rectangles, rect)

	// negative Y-index
	rect = common.NewDefaultRectangle()
	rect.Location.Y = -1
	rectangles = append(rectangles, rect)

	// negative X and Y-indexes
	rect = common.NewDefaultRectangle()
	rect.Location.X, rect.Location.Y = -1, -1
	rectangles = append(rectangles, rect)

	// empty height
	rect = common.NewDefaultRectangle()
	rect.Height = 0
	rectangles = append(rectangles, rect)

	// empty width
	rect = common.NewDefaultRectangle()
	rect.Width = 0
	rectangles = append(rectangles, rect)

	// outline and fill nil
	rect = common.NewDefaultRectangle()
	rect.Outline, rect.Fill = nil, nil
	rectangles = append(rectangles, rect)

	return rectangles
}
