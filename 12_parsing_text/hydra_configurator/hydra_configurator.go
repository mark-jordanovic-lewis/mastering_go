package hydra_configurator

import (
	"errors"
	"reflect"
)

const (
	CUST uint8 = iota
	)

var wrongTypeError error = errors.New("Type must be a pointer to a struct")

func GetCustomConfiguration(confType uint8, obj interface{}, filename string) (err error) {
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
	case CUST:
		err = MarshalCustomConfig(reflectedVal, filename)
	}
	return err
}