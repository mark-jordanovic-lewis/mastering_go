package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type CrewMember struct {    // tags define Marshaling behaviour (key, type, and if they are left out) lowercase fields are not exportable!
	ID 				  int  		`xml:"id,string,omitempty,attr"`
	Name              string	`xml:"name,omitempty"`
	SecurityClearance int		`xml:"clearance"`
	AccessCodes 	  []string 	`xml:"access>code,omitempty"`
}

type ShipInfo struct {
	XMLName      xml.Name	`xml:"Ship"`
	ShipID       int		`xml:"ident,string,attr"`
	ShipClass    string	    `xml:"Details>class"`
	NumberOfGuns int		`xml:"Details>n_guns"`
	Captain      CrewMember	`xml:",omitempty"`
}

func main()  {
	// Marshal/Unmarshal a string
	var member = CrewMember{ID:1, Name:"JJ", SecurityClearance:3,AccessCodes:[]string{"abc", "def"}}
	var ship = ShipInfo{ShipID:1, ShipClass:"Big", Captain:member, NumberOfGuns:2}

	var xmlBytes, err = xml.MarshalIndent(ship, " ","	")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("ship xml: \n", xml.Header, string(xmlBytes))
	var emptyShip = ShipInfo{}
	err = xml.Unmarshal(xmlBytes, &emptyShip)
	fmt.Printf("ship struct: %+v\n", emptyShip)


	// Encoding/Decoding a file
	file, err := os.Create("ship.xml")
	var xmlEncoder = xml.NewEncoder(file)
	xmlEncoder.Indent(" ", "	")
	err = xmlEncoder.Encode(ship)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	_ = file.Close()

	file, err = os.Open("ship.xml")
	emptyShip = ShipInfo{}
	var xmlDecoder = xml.NewDecoder(file)
	err = xmlDecoder.Decode(&emptyShip)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("ship struct read from file: %+v\n", emptyShip)
	_ = file.Close()
}

