package userstore

type RectangleOutOfBoundsError struct{}

func (e RectangleOutOfBoundsError) Error() string {
	return "rectangle index out of bounds"
}

type DrawingOutOfBoundsError struct{}

func (e DrawingOutOfBoundsError) Error() string {
	return "drawing index out of bounds"
}

type EmptyDrawingParamsError struct{}

func (e EmptyDrawingParamsError) Error() string {
	return "fill and outline fields are empty"
}
