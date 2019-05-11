package once

import (
	"sync"
	"fmt"
	)

// see also 08_design_patterns/singleton

func OneOutput() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("This should wait until the worker groups are complete and run only once.")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}