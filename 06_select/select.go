package main

import (
  "fmt"
  "math/rand"
  "time"
)

var clearanceMap = map[string]int {
  "Jenny": 5,
  "Kevin": 10,
  "CJ": 3,
}

func findSecClearance(name, server string, out chan int) {
  // Simulate Search
  time.Sleep(time.Duration(rand.Intn(10))* time.Second)
  // pipe into channel
  out <- clearanceMap[name]
}

func main() {
  defer func() {
    if err := recover(); err != nil {
      fmt.Println("The panic was recovered:\n", err)
    }
  }()
  rand.Seed(time.Now().UnixNano())

  chan_postgres := make(chan int)
  chan_mongo := make(chan int)
  name := "Jenny"
  go findSecClearance(name, "Postgres Server", chan_postgres)
  go findSecClearance(name, "Mongo Server", chan_mongo)

  db_search_loop:
  for {
    select {
    case clearance := <- chan_postgres:
      fmt.Println(name, "has a clearance of", clearance , "found in postgres")
      break db_search_loop
    case clearance := <- chan_mongo:
      fmt.Println(name, "has a clearance of", clearance, "found in mongo")
      break db_search_loop
    case <- time.After(time.Duration(3*time.Second)):
      panic(fmt.Sprintf("\nCould not find clearance of: %s\nBoth db timedout\n", name))
      // fmt.Errorf("\nCould not find clearance of: %s\nBoth db timedout\n", name)
      break db_search_loop
    default:
      time.Sleep(time.Duration(50*time.Millisecond))
      fmt.Print(".")
    }
  }
}
