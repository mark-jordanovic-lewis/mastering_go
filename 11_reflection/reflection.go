package main

// reflection casts from an interface instance to the concrete type.
// TypeOf and ValueOf help us with this.

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	var x1 float32 = 3.142
	x2, err := checkAndReturnFloatWithValues(x1)
	if err!= nil {
		fmt.Println("Was not passed float32")
	} else if x1 == x2 {
		fmt.Println("Managed to coerce the same value float.")
	}
	var floatInterface, err = generateInterfaceFromFloat32(x2)
	if err != nil {
		//handle err
	}
	// checking and setting values from reflections
	var reflectionOfX1 = reflect.ValueOf(x1)
	fmt.Println("the value is not settable as it is a reflection of a copied value: cN", reflectionOfX1.CanSet())

}

func checkAndReturnFloatWithValues(x interface{}) (float32, error)  {
	var reflectedValue = reflect.ValueOf(x)
	if reflectedValue.Kind() == reflect.Float32 {
		fmt.Printf("value struct of x: %+v\n", reflectedValue)
		fmt.Printf("Type of x: %+v\n", reflectedValue.Type())
		return float32(reflectedValue.Float()), nil
	} else {
		return 0, errors.New("Not a float32")
	}
}

func generateInterfaceFromFloat32(f float32) (interface{}, error) {
	var reflectedInterface = reflect.ValueOf(f).Interface()

	switch reflectedInterface.(type) {
	case float32:
		fmt.Println("found a float32, here's an interface of it")
		return reflectedInterface, nil
	default:
		return struct{}{}, errors.New("Did not find a float32 somehow")
	}
}

