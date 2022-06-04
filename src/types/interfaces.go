package types

import (
	"exercise/src/types/common"
)

type CanvasDrawing interface {
	FillCanvas(common.Rectangle, rune, rune) error
	GetCanvas() common.Canvas
}
