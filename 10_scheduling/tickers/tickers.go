package tickers

import (
	"fmt"
	"time"
)

func TickCounter(n int)  {
	var ticker = time.NewTicker(time.Duration(n)*time.Millisecond)
	var i = 0
	for t := range ticker.C {
		i++
		fmt.Printf("TickCounter - Tick number: %v, per %v millisceonds at %v\n", i, n, t)
	}
}

func PassedTickCounter(ticker *time.Ticker, complete chan bool)  {
	Loop:
	for {
		select {
		case <-complete:
			fmt.Println("Got stop signal")
			break Loop
		case t := <-ticker.C:
			fmt.Printf("PassedTickCounter - %v\n", t)
		}
	}
	fmt.Println("Exiting PassedTickCounter")
}