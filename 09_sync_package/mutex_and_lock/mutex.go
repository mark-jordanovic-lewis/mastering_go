package mutex_and_lock

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// RWMutex
// - read lock allows multiple read, disallows writing
// - write lock locks all
// - handy for concurrency safe maps
type SafeMap struct {
	notArray map[int]int
	sync.RWMutex
}

func (self *SafeMap) Initialise() {
	self.notArray = make(map[int]int)
}

func (self *SafeMap) WriteData(data []int)  {
	for index, input := range data {
		self.Lock()
		self.notArray[index] = input
		self.Unlock()
		time.Sleep(1*time.Second)
	}
}

func (self *SafeMap) ReadData(n int)  {
	for {
		self.RLock()
		mapSize := len(self.notArray)
		index := rand.Intn(mapSize)
		selection := self.notArray[index]
		self.RUnlock()
		fmt.Printf("READER#%v random selection: { index: %v, value: %v }\n", n, index, selection)
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	}
}

func (self *SafeMap) RemoveEvens() {
	for {
		self.RLock()
		mapSize := len(self.notArray)
		index := rand.Intn(mapSize)
		selection := self.notArray[index]
		if index%2 == 0 && selection != 0 {
			self.RUnlock()
			fmt.Printf("Zeroing: { index: %v, value: %v }\n", index, selection)
			self.setToZero(index)
		} else {
			self.RUnlock()
		}
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func (self *SafeMap) setToZero(index int)  {
	self.Lock()
	self.notArray[index] = 0
	self.Unlock()
}