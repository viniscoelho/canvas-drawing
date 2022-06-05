package canvas

type RectangleOutOfBoundsError struct{}

func (e RectangleOutOfBoundsError) Error() string {
	return "rectangle index out of bounds"
}

type DrawingOutOfBoundsError struct{}

func (e DrawingOutOfBoundsError) Error() string {
	return "drawing index out of bounds"
}
