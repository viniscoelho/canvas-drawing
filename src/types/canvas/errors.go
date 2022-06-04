package userstore

type RectangleOutOfBoundsError struct{}

func (e RectangleOutOfBoundsError) Error() string {
	return "Rectangle index out of bounds"
}

type DrawingOutOfBoundsError struct{}

func (e DrawingOutOfBoundsError) Error() string {
	return "Drawing index out of bounds"
}
