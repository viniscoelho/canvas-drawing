package common

type CanvasDTO struct {
	Canvas CanvasString `json:"canvas"`
}

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
