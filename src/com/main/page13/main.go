package main

import (
	"math"
	"fmt"
)

func main() {
	// 个人理解为,hypot成为他的函数名.
	// 函数也是值。Function values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}
