//go:generate mockgen -destination=mocks/mocks.go -package=mocks exercise/src/types CanvasDrawing
package types

import (
	"canvas-drawing/src/types/common"
)

type CanvasDrawing interface {
	FillCanvas(common.Rectangle) error
	GetCanvas() common.CanvasString
}
