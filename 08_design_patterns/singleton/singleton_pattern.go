package main

// this is a logger implementation

import (
	"./hlogger"
)

func main()  {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service mock logging")
	for i := 0; i < 10; i ++ {
		logger.Println(i)
	}
}