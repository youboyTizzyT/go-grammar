package myTest

import (
	"fmt"
	"sort"
	"sync"
	"testing"
	"time"
)
var wg = sync.WaitGroup{}
func TestThread1(t *testing.T) {
	for i := 0; i < 1; i++ {
		go forRun(int32(i))
	}
	time.Sleep(1*time.Second)
	wg.Wait()
}
func TestThread2(t *testing.T) {
	for i := 0; i < 2; i++ {
		go forRun(int32(i))
	}
	time.Sleep(1*time.Second)
	wg.Wait()
}
func TestThread3(t *testing.T) {
	for i := 0; i < 3; i++ {
		go forRun(int32(i))
	}
	time.Sleep(1*time.Second)
	wg.Wait()
}
func TestThread4(t *testing.T) {
	for i := 0; i < 4; i++ {
		go forRun(int32(i))
	}
	time.Sleep(1*time.Second)
	wg.Wait()
}
func TestThread8(t *testing.T) {
	for i := 0; i < 8; i++ {
		go forRun(int32(i))
	}
	time.Sleep(1*time.Second)
	wg.Wait()
}
func TestThread12(t *testing.T) {
	for i := 0; i < 12; i++ {
		go forRun(int32(i))
	}
	time.Sleep(1*time.Second)
	wg.Wait()
}
func TestThread16(t *testing.T) {
	for i := 0; i < 16; i++ {
		go forRun(int32(i))
	}
	time.Sleep(1*time.Second)
	wg.Wait()
}

func TestThread20(t *testing.T) {
	for i := 0; i < 200; i++ {
		go forRun(int32(i))
	}
	time.Sleep(1*time.Second)
	wg.Wait()
}
/*
for循环执行1000次
 */
func forRun(num int32) {
	wg.Add(1)
	defer wg.Done()
	var echoFunc =make(times,0)
	start:=time.Now()
	for i := 0; i < 1000; i++ {
		echoFunc=runFunc(echoFunc)
		time.Sleep(1*time.Millisecond)
	}
	fmt.Println("总共耗时",num,time.Now().Sub(start))
	aa:=time.Duration(0)
	for _, duration := range echoFunc {
		aa+=duration
	}
	sort.Sort(echoFunc)
	fmt.Println("平均延迟",num,aa/time.Duration(len(echoFunc)))
	fmt.Println("top 5 Max延迟",num,echoFunc[len(echoFunc)-5:])
	fmt.Println("top 5 Min延迟",num,echoFunc[:5])
}


//========================================================
/*
执行方法
 */
func runFunc(echoFunc []time.Duration) []time.Duration {
	start:=time.Now()
	feibonaqie(31)
	return append(echoFunc,time.Now().Sub(start) )
}
/*
测试算法
 */
func feibonaqie(int2 int) int {
	if int2 <= 0 {
		return -1
	}
	if int2==1 {
		return 1
	}
	if int2==2 {
		return 1
	}
	return feibonaqie(int2-1)+feibonaqie(int2-2)
}

type times []time.Duration

func (m times) Swap(i, j int)  {

	m[i], m[j] = m[j], m[i]
}

func (m times) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m times) Len() int {
	return len(m)
}
