package main

import (
  "fmt"
  "math/rand"
  "time"
)

type customRand struct {
  // Rand *rand.Rand // named fields in a struct are accessed as struct.Name.MethodName()
  *rand.Rand // unnamed struct methods are 'promoted' to be accessed directly: struct.MethodName()
  count int
}

func NewCustomRand(i int64) *customRand {
  return &customRand{
    Rand: rand.New(rand.NewSource(i)), // must initialise using the embedded Types typename if unnamed in the struct declaration
    count: 0,
  }
}

func (self *customRand) randRange(min, max int) int {
  self.count ++
  return self.Intn(max-min) + min
}

func (self *customRand) GetCount() int {
  return self.count
}

func (self *customRand) Intn(max int) int {
  fmt.Println("Called outer method")
  return self.Rand.Intn(max)+5
}

func main() {
  myRand := NewCustomRand(time.Now().UnixNano())
  fmt.Printf("rand between 5 and 40: %v", myRand.randRange(5,40))
  fmt.Printf("Some random float: %v", myRand.Float64())                // calling a promoted Rand method
  fmt.Printf("rand between 5 and 40: %v", myRand.Intn(35))       // this method call should call the CustomRand method
  fmt.Printf("rand between 5 and 40: %v", myRand.Rand.Intn(35))    // this will call the Rand method
  fmt.Printf("Number of time called: %v", myRand.GetCount())
}

