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

func Channels() {
	channel := make(chan string)
	go AddToChannel(1, channel)
	go AddToChannel(2, channel)
	go AddToChannel(3, channel)
	go AddToChannel(4, channel)
	go AddToChannel(5, channel)
	go AddToChannel(6, channel)
	go AddToChannel(7, channel)
	go AddToChannel(8, channel)
	go AddToChannel(9, channel)
	go AddToChannel(10, channel)
	println("here")
	go CloseChannel(channel)
	for data := range channel {
		println(data)
	}
}

func CloseChannel(channel chan string) {
	for {
		time.Sleep(1 * time.Second)
		if openRoutines == 0 {
			println("no more open routines, closing channel")
			close(channel)
			return
		}
	}
}

func AddToChannel(routineNumber int, channel chan string) {
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
