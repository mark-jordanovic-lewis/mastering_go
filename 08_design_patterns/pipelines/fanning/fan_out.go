package fanning

func FanOut(in <-chan int, outs ...chan<- int)  {
	for input := range in {
		go func() {
			broadcast(input, outs)
		}()
	}
}

func broadcast(input int, outs ...chan<- int)  {
	for _, channel := range outs {
		go func(channel chan<- int) {
			channel<- input
		}(channel)
	}
}