package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type task struct {
	id   string
	take time.Time
}

var c = make(chan *task, 10)

func main() {
	defluteNum := 0
	sleepNum := 0
	//CserverThreadWorkEmptyNum := 100000
	go func() {
		for true {
			select {
			case aa := <-c:
				if time.Since(aa.take) > time.Millisecond {
					fmt.Println(time.Now(), "take aaa ", aa.id, time.Since(aa.take))
				}
				defluteNum = 0
				//default:
				//	if defluteNum >= CserverThreadWorkEmptyNum {
				//		//time.Sleep(time.Duration(1) * time.Nanosecond)
				//		//runtime.Gosched()
				//		sleepNum++
				//	} else {
				//		defluteNum++
				//	}
			}
		}
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			i2 := rand.Int()

			c <- &task{id: strconv.Itoa(i2), take: time.Now()}
			//time.Sleep(time.Millisecond)
		}
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(sleepNum)
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Minute)
}
