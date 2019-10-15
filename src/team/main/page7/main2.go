package main

import "fmt"

/**
斐波那契数
*/
func fib() func() int {
	a, b := 0, 1
	c := func() int {
		a, b = b, a+b
		return a
	}
	return c
}

// 函数内部函数。每次初始化，将会初始化外部的函数内变量值
func main() {
	f := fib()
	fmt.Println(f(), f(), f(), f(), f(), f())
	d := fib()()
	fmt.Println(d, d, d, d, d, d)
	fmt.Println(fib()(), fib()(), fib()(), fib()(), fib()())
}
