package common

import (
	"errors"
	"fmt"
)

// TODO: should Outline and Fill be a string instead?
type Rectangle struct {
	Location Coordinates
	Width    int
	Height   int
	Outline  *rune
	Fill     *rune
}

type Coordinates struct {
	X int
	Y int
}

func NewDefaultRectangle() Rectangle {
	outline := '@'
	fill := 'X'
	return Rectangle{
		Location: Coordinates{
			X: 0,
			Y: 0,
		},
		Width:   3,
		Height:  3,
		Outline: &outline,
		Fill:    &fill,
	}
}

func NewRectangleFromDTO(rectDTO RectangleDTO) (Rectangle, error) {
	if err := validateRectangleDTO(rectDTO); err != nil {
		return Rectangle{}, fmt.Errorf("could not create rectangle from dto: %w", err)
	}

	rect := Rectangle{
		Location: Coordinates{
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

func validateRectangleDTO(rect RectangleDTO) error {
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
