package main

import (
	"fmt"
	"math"
)
/**
 * This is Printer!
 * 布尔值：false
 * 二进制：11111111
 * 八进制：377
 * 十六进制：FF
 * 十进制：255
 * 浮点数：3.141593
 * 字符串：printer
 *
 * 对象类型：int,string,bool,float64
 * 集合：[1 2 3 4 5]
 */
func main() {
	/**
	Golang加入了%T打印值的类型，%v打印数组等集合的所有元素。
	 */
	fmt.Println("This is Printer!")
	fmt.Printf("布尔值: \t%t\n",false)
	fmt.Printf("二进制: \t%b\n",255)
	fmt.Printf("八进制: \t%o\n",255)
	fmt.Printf("十六进制: \t%X\n",255)
	fmt.Printf("十进制: \t%b\n",255)
	fmt.Printf("浮点型: \t%b\n",math.Pi)
	fmt.Printf("字符串: \t%s\n","hello world!")
	fmt.Printf("对象类型: \t%T,%T,%T,%T\n",1,"hello",true,math.E)
	fmt.Printf("集合:   \t%v\n",[5]int{1,2,3,4,5})
}