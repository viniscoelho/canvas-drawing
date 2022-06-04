package types

import (
	"exercise/src/types/common"
)

type CanvasDrawing interface {
	Fill(common.Rectangle, rune, rune) error
	Print()
}
