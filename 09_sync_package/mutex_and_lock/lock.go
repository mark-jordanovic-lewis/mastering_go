package mutex_and_lock

import "sync"

type SafeCounter struct {
	i int
	sync.Mutex
}

func (self *SafeCounter) Increment() {
	self.Lock()
	self.i++
	self.Unlock()
}

func (self *SafeCounter) Decrement() {
	self.Lock()
	self.i--
	self.Unlock()
}

// lack of lock means that this can be called a the same time so potential race condition
// running `go run -race sync.go` will give us the warnings
func (self *SafeCounter) GetCounter() int  {
	return self.i
}

func (self *SafeCounter) GetSafeCounter() int  {
	self.Lock()
	count := self.i
	self.Unlock()
	return count
}