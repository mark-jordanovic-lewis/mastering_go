package main

import (
	h_conf "./hydra_configurator"
	"fmt"
)

// of course, the struct format must be known at compile time to read into
// and so must match the json contract
type ConfStruct struct {
	TS 		string  `name:"testString"`
	TB 		bool    `name:"testBool"`
	TF 		float64 `name:"testFloat"`
	TestInt int
}

func main() {
  var configStruct = new(ConfStruct)
  _ = h_conf.GetCustomConfiguration(h_conf.CUSTOM, configStruct, "config/config.conf")
  fmt.Printf("initialised configStruct: %+v\n", *configStruct)

  fmt.Printf("%t - %v\n",configStruct.TS,configStruct.TS)
  fmt.Printf("%t - %v\n",configStruct.TB,configStruct.TB)
  fmt.Printf("%t - %v\n",configStruct.TF,configStruct.TF)
  fmt.Printf("%t - %v\n",configStruct.TestInt,configStruct.TestInt)
}
