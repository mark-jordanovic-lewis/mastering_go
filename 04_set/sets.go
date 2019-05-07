package main

import "fmt"

type Set map[string]struct{}

func main() {
  set := make(Set)
  set["a"] = struct{}{} // empty struct wth empty values
  set["b"] = struct{}{}
  fmt.Println(setValues(set))
}

func setValues(set Set) (values []string) {
  for val, _ := range set { values = append(values, val) }
  return
}
