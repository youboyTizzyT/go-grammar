package _defer

import (
	"fmt"
)
/**
首先要明确的是：defer是在return之前执行的。这个在 官方文档中明确说明了的。

 */
func Defer() {
	fmt.Println(myDefer1())
	fmt.Print("==============\n")
	fmt.Println(myDefer2())
	fmt.Print("==============\n")
	fmt.Println(myDefer3())
	fmt.Print("==============\n")
	fmt.Println(*defertest1())
	}
/*
不懂
 */
func myDefer1()(result int) {
	result=2
	defer func() {
		result++
	}()
	return
}
func myDefer2() (r int) {
	t:=5
	defer func() {
		t=t+5
	}()
	return t
}
func myDefer3() (r int) {
	defer func(r int) {
		r=r+5
	}(r)
	return 1
}
func defertest1()*int{ //defer不会调用，仅仅会压入本地方法栈中，放在：｛所有栈变量之后，ret之前｝
	var i int
	defer func() {
		i++
		fmt.Println("c defer2:", i) // 打印结果为 c defer: 2
	}()
	defer func() {
		i++
		fmt.Println("c defer1:", i) // 打印结果为 c defer: 1
	}()
	return &i
}
