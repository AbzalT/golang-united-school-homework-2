package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type intCustomType int

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	var result float64
	sideLenSquare := sideLen * sideLen
	switch sidesNum {
	case SidesCircle:
		result = math.Pi * sideLenSquare
	case SidesTriangle:
		result = (math.Sqrt(3) * sideLenSquare) / 4
	case SidesSquare:
		result = sideLenSquare
	default:
		result = 0
	}

	return result
}
