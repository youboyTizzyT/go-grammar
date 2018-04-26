package main

import "fmt"

func fib() func() int {
	a, b := 0, 1
	c:=func() int {
		a, b = b, a+b
		return a
	}
	return c
}

func main() {
	f := fib()
	fmt.Println(f(), f(), f(), f(), f(),f())
	d:=fib()()
	fmt.Println(d,d,d,d,d,d)
	fmt.Println(fib()(),fib()(),fib()(),fib()(),fib()())
}
