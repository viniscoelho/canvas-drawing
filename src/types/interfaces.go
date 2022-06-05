package types

import (
	"exercise/src/types/common"
)

type CanvasDrawing interface {
	FillCanvas(common.Rectangle) error
	GetCanvas() common.CanvasString
}
