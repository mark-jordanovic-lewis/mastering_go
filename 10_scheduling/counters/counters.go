package counters

import (
	"fmt"
	"time"
)

func SlowCounter(n int)  {
	var i = 0
	var duration = time.Duration(n) * time.Millisecond

	for {
		var t = time.NewTimer(duration)
		<-t.C // this is the channel which the timer sends down every 'duration' that has been set
		i++
		fmt.Printf("SlowCounter: %v milliseconds have passed.\n", n * i)
	}
}

func SlowAfterCounter(n int)  {
	var i = 0
	var duration = time.Duration(n) * time.Millisecond

	for {
		<-time.After(duration)
		i++
		fmt.Printf("SlowAfterCounter: %v milliseconds have passed.\n", n * i)
	}
}

func ChanneledCounter(n int, durationChan chan int, stopChan chan bool)  {
	var i = 0
	var duration = time.Duration(n) * time.Millisecond
Loop:
	for {
		select {
		case <-time.After(duration):
			i++
			fmt.Printf("ChanneledCounter: %v milliseconds have passed.\n", n * i)
		case n = <-durationChan:
			fmt.Printf("ChanneledCounter: Got a new time period: %v milliseconds\n", n)
			duration = time.Duration(n)*time.Millisecond
		case <-stopChan:
			fmt.Println("Received stop signal")
			break Loop
		}
	}
	fmt.Println("Exiting ChanneledCounter")
}