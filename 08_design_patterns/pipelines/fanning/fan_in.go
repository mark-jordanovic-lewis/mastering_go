package fanning

func FanIn(channels ...<-chan int) <-chan int {
	var out = make(chan int)
	for _, channel := range channels {
		go registerChannel(channel, out)
	}
	return out
}

func registerChannel(in <-chan int, out chan<- int) {
	for input := range in {
		out <- input
	}
}