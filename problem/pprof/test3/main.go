package main

import (
	"net"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	// 网络模式
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "localhost:20010")
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	client, _ := net.DialTCP("tcp4", nil, tcpAddr)
	server, _ := listener.AcceptTCP()
	// chan模式
	var testChan = make(chan struct{}, 1024)

	go func() {
		http.ListenAndServe("localhost:20006", nil)
	}()
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
	// net server
	go func() {
		for true {
			server.Read([]byte{1})
			feibonaqie(20)

		}
	}()
	// net client
	go func() {
		t := time.NewTicker(time.Second / 30)
		defer t.Stop()
		for true {
			select {
			case <-t.C:
				client.Write([]byte{1})
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
