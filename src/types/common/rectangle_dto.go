package common

type RectangleDTO struct {
	Location *CoordinatesDTO `json:"coordinates"`
	Width    *int            `json:"width"`
	Height   *int            `json:"height"`
	Outline  *string         `json:"outline"`
	Fill     *string         `json:"fill"`
}

type CoordinatesDTO struct {
	X *int `json:"x"`
	Y *int `json:"y"`
}

func NewDTOFromRectangle(rect Rectangle) RectangleDTO {
	rectDTO := RectangleDTO{
		Location: &CoordinatesDTO{
			X: &rect.Location.X,
			Y: &rect.Location.Y,
		},
		Width:  &rect.Width,
		Height: &rect.Height,
	}

	if rect.Outline != nil {
		outline := string(*rect.Outline)
		rectDTO.Outline = &outline
	}

	if rect.Fill != nil {
		fill := string(*rect.Fill)
		rectDTO.Fill = &fill
	}

	return rectDTO
}
