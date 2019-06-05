package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type CrewMember struct {    // tags define Marshaling behaviour (key, type, and if they are left out) lowercase fields are not exportable!
	ID int                 	`json:"id,string"`
	Name string				`json:",omitempty"`
	SecurityClearance int
	AccessCodes []string 	`json:"Access,omitempty"`
	NickName string 		`json:"-"`
	nickName string
}

type ShipInfo struct {
	ShipID int				`json:"ident,string"`
	ShipClass string		`json:"class"`
	Captain CrewMember		`json:",omitempty"`
}

func main()  {
	// Marshal and Unmarshal structs to/from JSON strings
	var captain = CrewMember{1,"Jarvis", 10, []string{"beep", "boop"}, "R0b0Dude", "not exported"}
	var marshallable = map[string]interface{}{} // doesn't have to be a string key but this makes nice JSONs
	// don't have to do this step, just showing how powerful the marshaller is. These are different types being encoded to a json string held in a map
	marshallable["ship"] = ShipInfo{ShipID:1, ShipClass:"Actuiser", Captain: captain}
	marshallable["crew"] = []CrewMember{
		{2, "Bob", 3, []string{"none"}, "derrik", "dd"},
		{3, "Alice", 6, []string{"eighteen"}, "AlBoy", "aliz"},
		{5, "Billy", 2, []string{"mops"}, "stan", "soney"},
	}


	// Reading and decoding JSON from a map[string]interface{}{} - you have to know the types you are getting to begin with!
	// decoding and encodign only works if you have ExportedFields (starting with a capital letter!)
	type incomingStruct struct {
		Crew []CrewMember
		Ship ShipInfo
	}
	byteJson, err := json.Marshal(&marshallable) // can use both ref and value. Ref is cheaper tho.
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(string(byteJson))
	var unmarshallable = incomingStruct{}
	handleError(json.Unmarshal(byteJson, &unmarshallable))
	fmt.Printf("got unmarshallable:  %+v\n", unmarshallable)

	// write to Writable implementing type while encoding!
	jsonFile, err := os.Create("myShip.json")
	defer closeFile(jsonFile)
	handleError(json.NewEncoder(jsonFile).Encode(&marshallable))
	// read from Readable implementing type while Decoding
}

func handleError(err error)  {
	if err != nil {
		log.Fatal("Error occurred while processing JSON", err)
	}
}

func closeFile(f *os.File)  {
	err := f.Close()
	if err != nil {
		log.Println("Could not close file: ", f.Name())
	}
}