package hydra_configurator

import (
	"errors"
	"reflect"
)

const CUSTOM uint8 = iota

var wrongTypeError error = errors.New("Type must be a pointer to a struct")

func GetConfiguration(confType uint8, obj interface{}, filename string) (err error) {
	// check if have been passed a pointer
	var reflectedVal = reflect.ValueOf(obj)
	if reflectedVal.Kind() != reflect.Ptr || reflectedVal.IsNil() {
		return wrongTypeError
	}
	// get and confirm struct value
	reflectedVal = reflectedVal.Elem()
	if reflectedVal.Kind() != reflect.Struct {
		return wrongTypeError
	}

	switch confType {
	case CUSTOM:
		err = MarshalCustomConfig(reflectedVal, filename)
	}
	return err
}