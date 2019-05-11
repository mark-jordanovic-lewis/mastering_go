package main

import (
	"./mutex_and_lock"
	"./once"
	"fmt"
	"./wait_group"
	"sync"
	"time"
)
func main()  {
	// This runs a bunch of concurrent processes (may or may not be parallel)
	// using mutexes and wait groups to stagger and ensure the work has no race
	// conditions

	// lock.go usage
	// -------------
	sc := mutex_and_lock.SafeCounter{}
	// mutex wraps increment/decrement so should never be in locked state in output
	// as these are run in separate go routines the value does not switch between 0
	// and 1 but depends on which go routine manages to run first. ALways ends up at
	// zero as number of processes inc and dec are equal.
	for i := 0; i < 100; i++ {
		go sc.Increment()
		fmt.Printf("%v: %+v\n", i, sc)
		go sc.Decrement()
		fmt.Printf("%v: %+v\n", i, sc)
	}
	// race conditions happen here as the routines are pushed onto the stack and execution happens
	// during and after loop as spawning a routine is faster than executing lock-do_stuff-unlock so the
	// routines (potentially) continue running after the loop ends

	// Go _WILL_ run code with race conditions so need to be careful and run with -race while testing
	fmt.Println(sc.GetCounter()) // using GetSafeCounter() negates the rc

	// mutex.go usage
	// --------------
	sm := mutex_and_lock.SafeMap{}
	sm.Initialise()
	var data []int
	for i := 0; i < 100; i++ {
		data = append(data, i)
	}
	go sm.WriteData(data)
	time.Sleep(2*time.Second)
	go sm.ReadData(1)
	go sm.ReadData(2)
	go sm.RemoveEvens()
	time.Sleep(5*time.Second)

	// wait_group usage
	var group sync.WaitGroup
	wait_group.DoWork(&group)
	group.Wait()
	// once usage
	once.OneOutput()
	time.Sleep(2*time.Second)

}
