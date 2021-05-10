package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.Mutex

func main() {
	start := time.Now()
	ii := 0
	for i := 0; i < 100000000; i++ {
		wg.Lock() // todo 注释掉
		ii++
		wg.Unlock() // todo 注释掉
	}
	end := time.Now()
	fmt.Println(ii, end.Sub(start))
}
