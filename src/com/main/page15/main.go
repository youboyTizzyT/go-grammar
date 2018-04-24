package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	X, Y float64
}

// Go 没有类。然而，仍然可以在结构体类型上定义方法。
// 方法接收者 出现在 func 关键字和方法名之间的参数中。
func (v *Vertex) Abs() float64 {
	// 算术平方根
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// 事实上，可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。
// 不能对来自其他包的类型或基础类型定义方法。
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
