package main

import (
  "../08_design_patterns/singleton/hlogger"
  "fmt"
  "net/http"
)

func main() {
  logger := hlogger.GetInstance()
  logger.Println("Starting Hydra web service")

  http.HandleFunc("/", root)
  _ = http.ListenAndServe(":8080", nil)
}

func root(out http.ResponseWriter, in *http.Request) {
  logger := hlogger.GetInstance()
  _, _ = fmt.Fprintf(out, "Welcome to the hyrda software system")

  logger.Println("Received http on root url")
}
