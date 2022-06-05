package routes

import (
	"encoding/json"
	"errors"
	"exercise/src/types/common"
)

func newRectangleFromDTO(rectDTO common.RectangleDTO) (common.Rectangle, error) {
	if err := validateRectangle(rectDTO); err != nil {
		return common.Rectangle{}, err
	}

	rect := common.Rectangle{
		Location: common.Coordinates{
			X: *rectDTO.Location.X,
			Y: *rectDTO.Location.Y,
		},
		Width:  *rectDTO.Width,
		Height: *rectDTO.Height,
	}

	if rectDTO.Outline != nil {
		outline := rune((*rectDTO.Outline)[0])
		rect.Outline = &outline
	}
	if rectDTO.Fill != nil {
		fill := rune((*rectDTO.Fill)[0])
		rect.Fill = &fill
	}

	return rect, nil
}

func validateRectangle(rect common.RectangleDTO) error {
	if rect.Location == nil || rect.Location.X == nil || rect.Location.Y == nil {
		return errors.New("rectangle location cannot be nil")
	}

	if rect.Height == nil || rect.Width == nil {
		return errors.New("rectangle height and width cannot be nil")
	}

	if rect.Outline == nil && rect.Fill == nil {
		return errors.New("either outline or fill cannot be nil")
	}

	if rect.Outline != nil && len(*rect.Outline) != 1 {
		return errors.New("outline length must be at most one character long")
	}

	if rect.Fill != nil && len(*rect.Fill) != 1 {
		return errors.New("fill length must be at most one character long")
	}

	return nil
}

func serializeCanvas(canvas common.CanvasString) ([]byte, error) {
	dto := common.CanvasDTO{
		Canvas: canvas,
	}
	return json.Marshal(dto)
}
