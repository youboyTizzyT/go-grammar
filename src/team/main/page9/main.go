package main

import "fmt"
// 一个结构体（`struct`）就是一个字段的集合。

// 点结构体
type Point struct {
	X int
	Y int
}
// 线结构体
type Line struct {
	A Point
	B Point
}

func main() {
	p0:=Point{1,2}
	p1:=Point{3,4}
	linr:=Line{p0,p1}
	fmt.Println(linr)
	fmt.Println(p1,p0)
	// 结构体字段使用点号来访问。
	linr.A.Y=5
	fmt.Println(linr)
	q:=&p0
	q.X=3
	fmt.Println(q)
}
