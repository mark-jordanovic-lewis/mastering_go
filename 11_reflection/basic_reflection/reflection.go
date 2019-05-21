package basic_reflection

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
	floatInterface, err := generateInterfaceFromFloat32(x2)
	if err != nil {
		//handle err
	}
	fmt.Printf("value: %+v and type: %t of floatInterface.\n", floatInterface, floatInterface) // not sure how to check to see if the float is actually known as an interface and not strictly a float
	// checking and setting values from reflections
	fmt.Println("value of x1: ", x1)
	var reflectionOfX1 = reflect.ValueOf(x1)
	fmt.Println("  the value of x1 is not settable as a reflection of a copied value: can set - ", reflectionOfX1.CanSet())
	var reflectionOfRefToX1 = reflect.ValueOf(&x1)      // get reference to value
	fmt.Println("  the value of x1 is still not settable as a reflection of a reference to a value: can set - ", reflectionOfRefToX1.CanSet())
	var elementalValueOfx1 = reflectionOfRefToX1.Elem() // dereference value
	fmt.Println("  the value of x1 is now settable as an elemental reflection of a reference to a value: can set - ", elementalValueOfx1.CanSet())
	fmt.Printf("    current value of elementalValueOfx1: %+v\n", elementalValueOfx1)
	elementalValueOfx1.SetFloat(3.3)                // assign to memory location
	fmt.Printf("    new value of elementalValueOfx1: %+v\n", elementalValueOfx1)
	fmt.Println("    value of x1: ", x1)
	// reflect pointer -> elemental dereference -> assign to memory location
	// is the same concept as method passing to alter value.
	ChangeX(&x1)
}

func ChangeX(x *float32) { // are is reference to value
	*x = 5.5 // dereference and assign to memory location
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

