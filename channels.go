package main

import (
	"strconv"
	"sync"
	"time"
)

var (
	openRoutines int
	counter      sync.Mutex
)

func channels() {
	channel := make(chan string)
	go addToChannel(1, channel)
	go addToChannel(2, channel)
	go addToChannel(3, channel)
	go addToChannel(4, channel)
	go addToChannel(5, channel)
	go addToChannel(6, channel)
	go addToChannel(7, channel)
	go addToChannel(8, channel)
	go addToChannel(9, channel)
	go addToChannel(10, channel)
	println("here")
	go closeChannel(channel)
	for data := range channel {
		println(data)
	}
}

func closeChannel(channel chan string) {
	for {
		time.Sleep(1 * time.Second)
		if openRoutines == 0 {
			println("no more open routines, closing channel")
			close(channel)
			return
		}
	}
}

func addToChannel(routineNumber int, channel chan string) {
	counter.Lock()
	openRoutines++
	counter.Unlock()
	for i := 0; i < 100; i++ {
		channel <- time.Now().String() + "  " + strconv.Itoa(routineNumber) + ": " + strconv.Itoa(i)
		time.Sleep(time.Duration(routineNumber) * time.Millisecond)
	}
	counter.Lock()
	openRoutines--
	counter.Unlock()
}
