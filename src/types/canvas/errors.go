package canvas

type RectangleOutOfBoundsError struct{}

func (e RectangleOutOfBoundsError) Error() string {
	return "rectangle index out of bounds"
}

type DrawingOutOfBoundsError struct{}

func (e DrawingOutOfBoundsError) Error() string {
	return "drawing index out of bounds"
}

type InvalidCanvasError struct{}

func (e InvalidCanvasError) Error() string {
	return "invalid height and/or width"
}

type InvalidParamsError struct{}

func (e InvalidParamsError) Error() string {
	return "outline and fill cannot be nil"
}
