// Using magicStore and sll I can make a Ruby like array

package main

import (
  "fmt"
  "errors"
)

type magicStore struct {
  value interface{}
  name string
}

var ErrUninitializedStore = errors.New("Please initialise the store")

func (self *magicStore) Set(val interface{}) error {
  if self == nil { return ErrUninitializedStore }
  self.value = val
  return nil
}

func (self *magicStore) Get() (interface{}, error) {
  if self == nil { return -1, ErrUninitializedStore }
  return self.value, nil
}

func NewMagicStore(name string) *magicStore {
  return &magicStore{ value: "Initial Value", name: name }
}

func main() {
  printType(0)
  printType(float32(0.5))
  printType("0 or 0.5")
  printType(true)
  printType(0.5)
  printType(uint(0))

  store := NewMagicStore("Integer Store")
  store.Set(4)
  v, _ := store.Get()
  if val, ok := v.(int); ok {
    fmt.Printf("%T store is holding: %d\n", val, val)
  }

}

func printType(in interface{}) {
  switch in.(type) {
  case int:
    fmt.Println("got an int")
  case string:
    fmt.Println("got a string")
  case float32:
    fmt.Println("got a float32")
  default:
    fmt.Printf("Type not handled: %T\n", in)
  }
}

// i don't think this will work, it will just re-wrap the types in an interface
func caster(in interface{}) (func() interface{}) {
  switch in.(type) {
  case int:
    return func() interface{}{
      val, _ := in.(int)
      return val
    }
  case string:
    return func() interface{} {
      val, _ := in.(string)
      return val
    }
  case float32:
    return func() interface{} {
      val, _ := in.(float32)
      return val
    }
  default:
    return func() interface{} {
      return "Could not handle type"
    }
  }
}
