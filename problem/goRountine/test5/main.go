package main

/*
#define _GNU_SOURCE
#include <sched.h>
#include <pthread.h>

void lock_thread(int cpuid) {
       pthread_t tid;
       cpu_set_t cpuset;

       tid = pthread_self();
       CPU_ZERO(&cpuset);
       CPU_SET(cpuid, &cpuset);
   pthread_setaffinity_np(tid, sizeof(cpu_set_t), &cpuset);
}
*/
import "C"
import (
	"fmt"
	"runtime"
	"time"
)

/*
out
6.0859119s
6.0690638s      9ms
*/
var runNum = 4000
var feibonaqieNum = 30

func main() {
	start := time.Now()
	chanRUn(runNum)
	fmt.Println(time.Since(start))
	start = time.Now()
	//forRunchanRUn(runNum)
	fmt.Println(time.Since(start))
}

var chanRun = make(chan struct{}, 1024)

func chanRUn(num int) {
	go func() {
		for i := 0; i < num; i++ {
			chanRun <- struct{}{}
		}
		close(chanRun)
	}()
	runtime.LockOSThread()
	for true {
		select {
		case _, ok := <-chanRun:
			if !ok {
				return
			}
			fmt.Printf("CPU: %d\n", C.sched_getcpu())
			feibonaqie(feibonaqieNum)
		}
	}
}
func forRunchanRUn(num int) {
	for i := 0; i < num; i++ {
		fmt.Printf("CPU: %d\n", C.sched_getcpu())
		feibonaqie(feibonaqieNum)
	}
}

// 斐波那契数列
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
