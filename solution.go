package square

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type CustomIntType int

const (
	SidesTriangle CustomIntType = 3
	SidesSquare   CustomIntType = 4
	SidesCircle   CustomIntType = 0
)

func CalcSquare(sideLen float64, sidesNum CustomIntType) float64 {
	switch sidesNum {
	case 3:
		fmt.Println("Triangle")
		return (math.Sqrt(3) / 4) * math.Pow(sideLen, 2)
	case 4:
		fmt.Println("Square")
		return math.Pow(sideLen, 2)
	case 0:
		fmt.Println("Circle")
		return math.Pi * math.Pow(sideLen, 2)
	default:
		fmt.Println("Not in list")
		return 0
	}

}
