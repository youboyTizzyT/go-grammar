package main

import (
	"fmt"
	"math"
)

func main() {
	// 个人理解为,hypot成为他的函数名. 其实，hypot是一个指针，指向堆里的函数体。
	// 函数也是值。Function values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}
