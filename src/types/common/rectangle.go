package common

type Canvas [][]rune

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
