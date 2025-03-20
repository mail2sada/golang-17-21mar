package advanced

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("Init from adv")
}
func Square(n int) int {
	return n * n
}

func Sin(angle float64) float64 {
	return math.Sin(angle)
}
