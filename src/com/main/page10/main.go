package main

import (
	"fmt"
	//"go/types"
)

func main() {
	// slice, 切片
	// 一个 slice 会指向一个数组，并且包含了长度信息。
	// []T 是一个元素类型为 T 的 slice。

	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n",
			i, p[i])
	}
	// slice 可以重新切片，创建一个新的 slice 值指向相同的数组。
	// 表达式

	// s[lo:hi]
	// 表示从 lo 到 hi-1 的 slice 元素，含两端。因此

	// s[lo:lo]
	// 是空的，而

	// s[lo:lo+1]
	// 有一个元素。
	fmt.Println("p[1:4] ==", p[1:4])
	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])
	// slice 由函数 make 创建。
	// 这会分配一个零长度的数组并且返回一个 slice 指向这个数组：
	// a := make([]int, 5)  // len(a)=5
	// 为了指定容量，可传递第三个参数到 `make`：
	// b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	// b = b[:cap(b)] // len(b)=5, cap(b)=5
	// b = b[1:]      // len(b)=4, cap(b)=4

	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	c[1]=5
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
	e := c[0:5]
	printSlice("e", e)
	// new的内置函数 它只接受一个参数，这个参数是一个类型，分配好内存后，返回一个指向该类型内存地址的指针。同时请注意它同时把分配的内存置为零，也就是类型的零值。
	var i=new(int)
	*i=10
	fmt.Print(*i)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v \n",s,len(s),cap(x),x)
}
