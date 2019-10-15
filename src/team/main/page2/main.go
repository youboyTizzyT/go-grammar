package main

// 导入
import (
	"fmt"
	"math"
)

func main() {
	// go语言也支持printf
	fmt.Printf("Now you have %g problems.\n", math.Nextafter(2, 3))
	// 首字母大写的成员才可以导出(即可以访问)
	// 改为pi将报错
	fmt.Println(math.Pi)
	fmt.Print("\n")
	fmt.Print(add(45, 84))
	a, b := swap("hello", "world")
	fmt.Print("\n")
	fmt.Println(a, b)
	fmt.Print("\n")
	fmt.Println(split(17))
}

// 定义函数
// 注意点: 返回类型在方法之后
// 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
func add(x, y int) int {
	// 调用函数
	return x + y
}

// 定义函数
// 函数可以返回任意数量的返回值。
func swap(x, y string) (string, string) {
	return y, x
}

// 命名返回值
// 函数接受参数。在 Go 中，函数可以返回多个“结果参数”，而不仅仅是一个值。它们可以像变量那样命名和使用。
// 如果命名了返回值参数，一个没有参数的 return 语句，会将当前的值作为返回值返回。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
