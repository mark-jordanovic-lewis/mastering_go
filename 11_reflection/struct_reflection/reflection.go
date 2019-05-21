package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"errors"
	"syscall"
)

type Printer interface {
	Print(s string)
}

type MyStruct struct {
	// no space between
	Field1 int `alias:"f1" desc:"first field of the struct"`
	Field2 string `alias:"f2" desc:"second field of the struct"`
	Field3 float64 `alias:"f3" desc:"third field of the struct"`
	field4 string `alias:"f4" desc:"fourth field of the struct"`
}

func (self *MyStruct) Print(s string) {
	self.field4 = s
	fmt.Println("I now say: ", self.field4)
	fmt.Println("F1:", self.Field1, "F2:", self.Field2, "F3:", self.Field3)
}

func main()  {
	var mystruct = MyStruct{2,"hello", 2.4, "I will break the compilation without the Print check"}
	InspectStruct(mystruct)
	err := UpdateStringField(&mystruct, 1, "I was updated") // updating requires reference!
	if err != nil {
		fmt.Println(err)
		syscall.Exit(1)
	}
	InspectStruct(mystruct)
	if ImplementsPrinter(&mystruct) { // look, passing as a reference
		callPrintOn(&mystruct)
	}
}

func InspectStruct(unknown interface{})  {
	var reflectedValue = reflect.ValueOf(unknown)
	var reflectedType = reflect.TypeOf(unknown)

	for i := 0; i < reflectedValue.NumField(); i++ {
		var fieldType= reflectedType.Field(i)
		var fieldVal= reflectedValue.Field(i)
		// .Interface() will not work on unexported fields, ie/ you cannot make them directly accessible by reflection.
		fmt.Printf("Reflected Field #%v:\n", i)

		// checking for accessible field is a bit of a ballache
		lowercaseRegex, _ := regexp.Compile("[a-z]")
		inaccessibleField := lowercaseRegex.MatchString(strings.Split(fieldType.Name, "")[0])

		if inaccessibleField {
			fmt.Printf("  fieldType (FieldStruct): %+v\n  fieldVal:  %t - %+v\n", fieldType, fieldVal, fieldVal)
		} else {
			fmt.Printf("  fieldType (FieldStruct): %+v\n  fieldVal:  %t - %+v\n", fieldType, fieldVal.Interface(), fieldVal)
		}
		fmt.Printf("  Name: %v\n  Tag: %v\n  Byte Offset of the field: %v\n  Field Index: %v\n  Anonimity of the field: %v\n", fieldType.Name, fieldType.Tag, fieldType.Offset, fieldType.Index, fieldType.Anonymous)
		fmt.Println("    Tags - alias: ", fieldType.Tag.Get("alias"), ", description: ", fieldType.Tag.Get("desc"))
	}
}



func ImplementsPrinter(unknown interface{}) bool {
	var val = reflect.ValueOf(unknown)
	var typ = val.Type()
	var thisInterface = reflect.TypeOf((*Printer)(nil)).Elem()
	var implements = typ.Implements(thisInterface)
	fmt.Println("Unknown interface implements Printer? - ", implements)
	return implements
}



func callPrintOn(unknown interface{}) {
	var val = reflect.ValueOf(unknown)
	var printFunc = val.MethodByName("Print")
	var args = []reflect.Value{ reflect.ValueOf("this new string") }
	printFunc.Call(args)
}



func UpdateStringField(unknown interface{}, index int, updateString string) error {
	var reflectedValue = reflect.ValueOf(unknown)
	if reflectedValue.Kind() != reflect.Ptr {
		return errors.New("cannot update non-pointer reflection")
	}
	if !canUpdateStringField(reflectedValue, index) {
		return errors.New("cannot update the field, no string at that index")
	}
	reflectedValue.Elem().Field(index).SetString(updateString)
	return nil
}



func canUpdateStringField(reflection reflect.Value, index int) bool {
	return reflection.Elem().NumField() >= index &&
		reflection.Elem().Field(index).Type().Name() == "string"
}