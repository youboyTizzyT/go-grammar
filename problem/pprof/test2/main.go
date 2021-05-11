package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:20006", nil)
	}()
	var testChan = make(chan struct{}, 1024)
	go func() {
		for true {
			select {
			case <-testChan:
				feibonaqie(20)
			}
		}
	}()
	go func() {
		t := time.NewTicker(time.Second / 30)
		defer t.Stop()
		for true {
			select {
			case <-t.C:
				testChan <- struct{}{}
			}
		}
	}()
	time.Sleep(1 * time.Hour)
}

func feibonaqie(int2 int) int {
	if int2 == 0 {
		return 1
	}
	if int2 == 1 {
		return 1
	}
	if int2 == 2 {
		return 1
	}
	return feibonaqie(int2-1) + feibonaqie(int2-2)
}
