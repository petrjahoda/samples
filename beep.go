package main

import (
	"github.com/gen2brain/beeep"
	"time"
)

func beep() {
	i := 0
	for i < 3 {
		err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		if err != nil {
			panic(err)
		}
		time.Sleep(100 * time.Millisecond)
		i++
	}
}
