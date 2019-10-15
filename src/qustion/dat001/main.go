/**
 * @program: go-grammar
 *
 * @description:
 *
 * @author: Mr.weicong
 *
 * @create: 2019-10-15 16:02
 **/
package main

import (
	"fmt"
	"time"
)

/**
对于无缓冲的 channel，用 select 来处理通道超时时间，如果在接收超时后，退出接收，那么该通道的发送会被阻塞吗？

会阻塞吗？实现原理是什么？
*/

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "this is send"
		fmt.Println("send ok")
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
	time.Sleep(2 * time.Second)
	fmt.Println("over")
}
