package main

import (
	"fmt"
	"runtime"
	"time"
)

///*
//#define _GNU_SOURCE
//#include <sched.h>
//*/
//import "C"

func worker(id int) {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		//fmt.Printf("worker: %d, CPU: %d\n", id, C.sched_getcpu())
		//feibonaqie(3)
		//feibonaqie(20)
		//feibonaqie(31)
		//feibonaqie(32)
		feibonaqie(33)
	}
	fmt.Printf("worker: %d, costTime: %v\n", id, time.Now().Sub(start))
}

func feibonaqie(int2 int) int {
	if int2 <= 0 {
		return -1
	}
	if int2 == 1 {
		return 1
	}
	if int2 == 2 {
		return 1
	}
	return feibonaqie(int2-1) + feibonaqie(int2-2)
}

func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 2; i++ {
		go worker(1)
	}
	time.Sleep(20 * time.Second)
}
