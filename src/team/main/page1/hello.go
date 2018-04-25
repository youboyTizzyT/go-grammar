package main

import (
    "fmt"
    "os"
    "net"
    "time"
    "math/rand"
)

func main() {
	// 打印文字
    fmt.Println("Hello, 世界")
    // 打印时间
    fmt.Println("The time is", time.Now())
    // 打印文件
    fmt.Println("And if you try to open a file:")
    fmt.Println(os.Open("main.go"))
    // 打印网络
    fmt.Println("Or access the network:")
    fmt.Println(net.Dial("tcp", "www.google.team:80"))

    // 打印随机数, 没有设置种子,所以一直都是1
    fmt.Println("My favorite number is", rand.Intn(10))
    // 设置种子 根据unix的nano纳秒作为随机种子
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    // 这样就能拿到随机值了
    fmt.Println("My favorite number is", r.Intn(10))
}
