package main

import (
	"./sheild_builder"
	"fmt"
)

func main()  {
	builder := sheild_builder.NewSheildBuilder()
	sheild := builder.RaiseFront().RaiseBack().Build()
	fmt.Printf("built sheild one: %+v\n", sheild)

	builder = sheild_builder.NewSheildBuilder()
	sheild = builder.
		RaiseFront().
		RaiseBack().
		RaiseLeft().
		RaiseRight().
		Build()
	fmt.Printf("built sheild two: %+v\n", sheild)
}