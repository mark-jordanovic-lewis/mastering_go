package wait_group

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Passing by reference as pass by value will clone
func DoWork(group *sync.WaitGroup)  {
	for i := 0; i <= 5; i++ {
		group.Add(1)
		go func(i int) {
			defer group.Done()
			workLoad := rand.Intn(1000)
			time.Sleep(time.Duration(workLoad)*time.Millisecond)
			fmt.Printf("%v milliseconds of work done for process #%v.\n", workLoad, i)
		}(i)
	}
}