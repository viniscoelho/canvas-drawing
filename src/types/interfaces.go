package types

import (
	"exercise/src/types/common"
)

type CanvasDrawing interface {
	DrawCanvas(common.Rectangle) error
	GetCanvas() common.CanvasString
}
