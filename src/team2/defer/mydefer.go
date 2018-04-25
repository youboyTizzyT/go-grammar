package _defer

import "fmt"
/**
首先要明确的是：defer是在return之前执行的。这个在 官方文档中明确说明了的。


 */
func Defer() {
	fmt.Println(myDefer1())
	fmt.Print("==============\n")
	fmt.Println(myDefer1())
	fmt.Print("==============\n")
	fmt.Println(myDefer3())
}
func myDefer1()(result int) {
	defer func() {
		result++
	}()
	return 0
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