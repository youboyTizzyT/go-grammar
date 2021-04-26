package myTest

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
	"testing"
	"time"
)

var wg = sync.WaitGroup{}

func BenchmarkThread1(t *testing.B) {
	forMain(t, 1)
}
func BenchmarkThread2(t *testing.B) {
	forMain(t, 2)
}
func BenchmarkThread4(t *testing.B) {
	forMain(t, 4)
}
func BenchmarkThread6(t *testing.B) {
	forMain(t, 6)
}
func BenchmarkThread8(t *testing.B) {
	forMain(t, 8)
}
func BenchmarkThread12(t *testing.B) {
	forMain(t, 12)
}
func BenchmarkThread16(t *testing.B) {
	forMain(t, 16)
}

func TestMain(m *testing.M) {
	runtime.GOMAXPROCS(16)
	currentcpu := runtime.GOMAXPROCS(-1)
	maxcpu := runtime.NumCPU()
	runtime.NumCPU()
	fmt.Println("开始", currentcpu, maxcpu)
	m.Run()
}
func forMain(t *testing.B, int2 int) {
	for i := 0; i < int2; i++ {
		go forRun(int32(1000))
	}
	time.Sleep(1 * time.Second)
	wg.Wait()
}

/*
for循环执行1000次
*/
func forRun(num int32) {
	wg.Add(1)
	defer wg.Done()
	var echoFunc = make(times, 0)
	//start := time.Now()
	for i := 0; i < int(num); i++ {
		runFunc(echoFunc)
		//echoFunc = runFunc(echoFunc)
		//time.Sleep(1 * time.Millisecond)
	}
	//fmt.Println("总共耗时", num, time.Now().Sub(start))
	aa := time.Duration(0)
	for _, duration := range echoFunc {
		aa += duration
	}
	sort.Sort(echoFunc)
	//fmt.Println("平均延迟", num, aa/time.Duration(len(echoFunc)))
	//fmt.Println("top 5 Max延迟", num, echoFunc[len(echoFunc)-5:])
	//fmt.Println("top 5 Min延迟", num, echoFunc[:5])
}

//========================================================
/*
执行方法
*/
func runFunc(echoFunc []time.Duration) {
	//start := time.Now()
	feibonaqie(31)
	//return
}

/*
测试算法
*/
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

type times []time.Duration

func (m times) Swap(i, j int) {

	m[i], m[j] = m[j], m[i]
}

func (m times) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m times) Len() int {
	return len(m)
}
