package common

type Canvas [][]string

type Rectangle struct {
	Location Coordinates
	Width    int
	Height   int
}

type Coordinates struct {
	X int
	Y int
}
