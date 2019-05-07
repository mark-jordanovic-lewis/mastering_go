package main

import (
  "net/http"
  "fmt"
)

func main() {
  http.HandleFunc("/", root)
  http.ListenAndServe(":8080", nil)
}

func root(out http.ResponseWriter, in *http.Request) {
  // write to out io.Writer (ResponseWriter)
  fmt.Fprintf(out, "Welcome to the hyrda software system")
}
