package main

import (
	h_conf "../../hydra_configurator"
	"fmt"
	"github.com/mark-jordanovic-lewis/mastering_go/12_parsing_text/hydra_configurator"
)

// of course, the struct format must be known at compile time to read into
// and so must match the json contract
type ConfStruct Struct {
	TS 		string  `name:"testString"`
	TB 		bool    `name:"testBool"`
	TF 		float64 `name:"testFloat"`
	TestInt int
}

func main() {
  var configStruct = new(ConfStruct)
  h_conf.GetConfiguration(h_conf.CUSTOM, configStruct, "config.conf")
  fmt.Printf("initialised configStruct: %+v\n", *configStruct)

  fmt.Printf("%t - %v\n",configStruct.TS,configStruct.TS)
  fmt.Printf("%t - %v\n",configStruct.TB,configStruct.TB)
  fmt.Printf("%t - %v\n",configStruct.TF,configStruct.TF)
  fmt.Printf("%t - %v\n",configStruct.TestInt,configStruct.TestInt)
}
