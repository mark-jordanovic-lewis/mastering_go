package main

import (
	h_conf "../hydra_configurator"
	"fmt"
)

type Conf struct {
	TS string  `name:"testString" xml:"testString" json:"testString"`
	TB bool    `name:"testBool" xml:"testBool" json:"testBool"`
	TF float64 `name:"testFloat" xml:"testFloat" json:"testFloat"`
	TI int32   `name:"testInt" xml:"testInt" json:"testInt"`
}

func main() {
	// passing by reference as writing to memory of the COPIED value further down the chain means nothing.
	var decodedConfig = Conf{}
	_ = h_conf.GetConfiguration(h_conf.JSON, &decodedConfig, "./config.json")
	fmt.Printf("decoded from JSON file: %+v\n", decodedConfig)

	decodedConfig = Conf{}
	_ = h_conf.GetConfiguration(h_conf.XML, &decodedConfig, "./config.xml")
	fmt.Printf("decoded from XML file: %+v\n", decodedConfig)


	// NB:
	// the new keyword returns a pointer to the new struct
	// - I prefer to be explicit and to have the value owned by the instantiating scope rather than have
	//   the runtime deal with scoping.
	// eg:
	// var newDecodedJSON = new(ConfJSON)
	// _ = h_conf.GetConfiguration(h_conf.JSON, newDecodedJSON, "./config.json")
	// fmt.Printf("decoded JSON using new instead of literal: %+v\n", newDecodedJSON)


}