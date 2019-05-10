package main

import (
	"fmt"
	"time"
)

// functions which returns a channel to communicate with the wider program

func main()  {
	var ticker = time.NewTicker(time.Duration(100)*time.Millisecond)
	var completeChannel = ChannelGenerator(ticker)
	time.Sleep(1*time.Second)
	ticker.Stop()
	completeChannel<-true
	close(completeChannel)

}

func ChannelGenerator(ticker *time.Ticker) chan bool  {
	var complete = make(chan bool)
	go func() {
		Loop:
			for {
				select {
				case t := <-ticker.C:
					fmt.Println("TICK AT: %v", t)
				case <-complete:
					fmt.Println("Got complete signal")
					break Loop
				}
			}
		fmt.Println("Exiting ChannelGenerator")
	}()
	return complete
}