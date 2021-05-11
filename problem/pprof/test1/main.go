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
	time.Sleep(1 * time.Hour)
}
