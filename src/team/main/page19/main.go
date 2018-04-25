package main

import (
	"time"
	"fmt"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000*time.Millisecond)
		fmt.Println(s,time.Now())
	}
}
func main() {
	go say("world")
	//say("hello")
	for true {
		i:=1
		i++
	}
}
