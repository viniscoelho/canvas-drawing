package canvas

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"exercise/src/types/common"
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
	cd.FillCanvas(rectA)

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
	cd.FillCanvas(rectB)

	expected := []string{
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
	for y := 0; y < len(expected); y++ {
		canvasRow := strings.TrimRight(canvas[y], " ")
		assert.Equal(expected[y], canvasRow, "canvas does not match")
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
	cd.FillCanvas(rectA)

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
	cd.FillCanvas(rectB)

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
	cd.FillCanvas(rectC)

	expected := []string{
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
	for y := 0; y < len(expected); y++ {
		canvasRow := strings.TrimRight(canvas[y], " ")
		assert.Equal(expected[y], canvasRow, "canvas does not match")
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
	cd.FillCanvas(rectA)

	expected := []string{
		".",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}

	canvas := cd.GetCanvas()
	for y := 0; y < len(expected); y++ {
		canvasRow := strings.TrimRight(canvas[y], " ")
		assert.Equal(expected[y], canvasRow, "canvas does not match")
	}
}
