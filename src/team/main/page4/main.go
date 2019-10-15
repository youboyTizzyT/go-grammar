package main

import "fmt"

func main() {

	// 变量定义可以包含初始值，每个变量对应一个。
	// 如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。
	// 在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。
	// 函数外的每个语法块都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。
	// go 可以自动判断类型
	//var x, y, z int = 1, 2, 3
	var x, y, z = 1, 2, 3.5
	c, python, java := true, false, "no!"
	fmt.Println(x, y, z, c, python, java)
}
