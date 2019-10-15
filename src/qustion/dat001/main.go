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
		c1 <- "this is send" // 此处不会执行成功。因为chan已经没有对应的接受的人了
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

/**
对于无缓冲的channel，发送时会阻塞在发送位置直到被接收后程序继续往下执行，接收时会阻塞在接收位置直到接收到信息然后程序继续往下执行。
是否阻塞和时间无关，是为了某种情景下认为的控制。就本例而言，“this is send"写入channel中，此时select模块已经执行完成，消息没有被接收，
因此"send ok"直到程序结束也不会被输出。
*/
