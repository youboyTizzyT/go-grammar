package main

import "fmt"

func main() {
	sum:=0
	// Go 只有一种循环结构——`for` 循环。
	// 没有`( )`
	// 必须有`{ }`
	for i := 0; i < 10; i++ {
		sum+=i
	}
	fmt.Println(sum)

	// 跟 C 或者 Java 中一样，可以让前置、后置语句为空。
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	// 基于此可以省略分号：C 的 while 在 Go 中叫做 `for`。
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// 如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环。
	var n =1
	for {
		fmt.Println(n)
		n++;
		if(n>10){
			break
		}
	}
}