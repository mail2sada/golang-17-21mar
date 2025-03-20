package geometry

const PI = 3.14159

func GetArea(radius float64) float64 {
	return PI * radius * radius
}

func GetDiameter(radius float64) float64 {
	return 2 * radius
}
