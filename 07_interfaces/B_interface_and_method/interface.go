package main

import (
  "fmt"
  "math/rand"
  "time"
)

type customRand struct {
  Rand *rand.Rand
  count int
}

func NewCustomRand(i int64) *customRand {
  return &customRand{
    Rand: rand.New(rand.NewSource(i)),
    count: 0,
  }
}

func (self *customRand) randRange(min, max int) int {
  self.count ++
  return self.Rand.Intn(max-min) + min
}

func main() {
  myRand := NewCustomRand(time.Now().UnixNano())
  fmt.Printf("rand between 5 and 40: %v", myRand.randRange(5,40))
  fmt.Printf("rand between 5 and 40: %v", myRand.Rand.Intn(35)+5)
  fmt.Printf("Number of time called: %v", myRand.count)
}
