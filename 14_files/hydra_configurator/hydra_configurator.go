package hydra_configurator

import (
	. "../../12_parsing_text/hydra_configurator"
	"errors"
	"reflect"
)

const (
	CUSTOM uint8 = iota
	JSON
	XML
	)

var wrongTypeError error = errors.New("Type must be a pointer to a struct")

// NB: The struct to be decoded INTO must be passed as a reference to this function here or the memory of the COPY is written to.
func GetConfiguration(confType uint8, obj interface{}, filename string) (err error) {
	// check if have been passed a pointer
	var reflectedVal = reflect.ValueOf(obj)
	if reflectedVal.Kind() != reflect.Ptr || reflectedVal.IsNil() {
		return wrongTypeError
	}
	// get and confirm struct value
	reflectedVal = reflectedVal.Elem() // Elem gets the value from the pointer, which is mutable
	if reflectedVal.Kind() != reflect.Struct {
		return wrongTypeError
	}

	switch confType {
	case CUSTOM:
		err = MarshalCustomConfig(reflectedVal, filename)
	case JSON:
		err = decodeJSONConfigFile(obj, filename)
	case XML:
		err = decodeXMLConfigFile(obj, filename)
	}
	return err
}