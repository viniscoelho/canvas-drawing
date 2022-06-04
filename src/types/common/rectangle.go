package common

type Rectangle struct {
	Location Coordinates
	Width    int
	Height   int
}

type Coordinates struct {
	Row    int
	Column int
}
