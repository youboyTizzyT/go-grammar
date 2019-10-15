package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")

	// switch 可以用于string
	// 除非使用 fallthrough 语句作为结尾，否则 case 部分会自动终止。
	// switch 可以在指向之前执行一段话
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	fmt.Println("When's Saturday?")
	// Weekday函数返回一个int值,但是执行打印的时候他会在time类中的一个days中去找对应的字符串,给人以假象
	today := time.Now().Weekday()
	fmt.Println(today)
	// switch 的条件从上到下的执行，当匹配成功的时候停止。 java中必须执行break才能停止
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()

	fmt.Println(t)
	// 没有条件的 switch 同 `switch true` 一样。
	// 这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening！！！")
	}
}
