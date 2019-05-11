package main

import (
	"./counters"
	"./tickers"
	"time"
)
// Timer: Unconsumed signals are immediately flushed causing specific increments in timing
//
// Ticker: Stores unconsumed signal so immediate consumption of latest signal waiting on channel
//
// Stop method will halt a timer/ticker signal send but will not close channel.
// Any coordination between go routines using the timer or spawned by the timer needs to be explicit
// The channels can be closed in the normal way. close(timer.C)

func main() {
	// Timers
	// ------
	go counters.SlowCounter(700)
	time.Sleep(2*time.Second)

	go counters.SlowAfterCounter(700)
	time.Sleep(2*time.Second)

	var durationChan = make(chan int)
	var stopChan = make(chan bool)
	go counters.ChanneledCounter(500, durationChan, stopChan)
	time.Sleep(1*time.Second)
	durationChan<- 300
	time.Sleep(2*time.Second)
	stopChan<-true
	time.Sleep(1*time.Second)

	// Tickers
	// =======
	go tickers.TickCounter(300)
	time.Sleep(2*time.Second)
	var ticker = time.NewTicker(time.Duration(200)*time.Millisecond)
	var complete = make(chan bool)
	go tickers.PassedTickCounter(ticker, complete)
	time.Sleep(2*time.Second)
	ticker.Stop()
	complete<-true
}
