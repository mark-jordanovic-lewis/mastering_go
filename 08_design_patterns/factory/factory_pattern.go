package main

import (
	"./appliances"
	"fmt"
)


func main()  {
	fmt.Println(" Which appliance do you wish to use?")
	fmt.Println("0: Stove")
	fmt.Println("1: Fridge")
	fmt.Println("2: Microwave")

	var myType int
	fmt.Scan(&myType)

	var myAppliance, err = appliances.CreateAppliance(myType)
	if err == nil {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		fmt.Println(err)
	}
}

