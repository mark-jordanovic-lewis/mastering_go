package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main(){
  response, err := http.Get("http://api.theysaidso.com/qod.json")
  if err != nil { fmt.Errorf("Could not GET `api.theysaidso.com/qod`: %s", err) }

  defer response.Body.Close()
  contents, err := ioutil.ReadAll(response.Body)
  if err != nil { fmt.Errorf("Could not parse response.Body: %+v\n%s", response.Body, err) }

  // casting to string as read data from all buffers is []rune == []int32
  fmt.Println("Body Content\n", string(contents))
}
