package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type CrewMember struct {
	ID 				  int
	Name              string
	SecurityClearance int
	AccessCodes 	  []string
}
type CrewMembers []CrewMember

func (self CrewMembers) toCsv() (rows [][]string) {
	for _, row := range self {
		var (
			id = strconv.Itoa(row.ID)
			clearance = strconv.Itoa(row.SecurityClearance)
			codes = strings.Join(row.AccessCodes, ",")
		)
		rows = append(rows, []string{id,row.Name,clearance,codes})
	}
	return
}

type ShipInfo struct {
	ShipID       int
	ShipClass    string
	NumberOfGuns int
	Captain      CrewMember
}

func main()  {
	file, err := os.Open("../assets/crew.csv")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer func() { _ = file.Close() }()

	var (
		csvReader  = csv.NewReader(file)
		csvSlice  CrewMembers
	)

	csvReader.Comment = '#'
	csvReader.Comma = ';'


	// Reading the CSV does require knowing the structure of the CSV before hand
	// so that particular fields can be parsed correctly
	// - could be twinned with a descriptor file like a csv which would help direct the operations
	var id = 0
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			// should record this in a log file so erroneous paring can be investigated.
			if pe, ok := err.(*csv.ParseError); ok {
				fmt.Println("bad column: ", pe.Column)
				fmt.Println("bad line: ", pe.Line)
				fmt.Println("Error: ", pe.Err)
				if pe.Err == csv.ErrFieldCount {
					continue
				}
			}
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("CSV row: ", record)
		var member = CrewMember{}
		member.Name = record[0]
		member.SecurityClearance, err = strconv.Atoi(record[1])
		member.AccessCodes = strings.Split(record[2], ",")
		if err != nil {
			fmt.Printf("%s has invalid clearance, skipping.", member.Name)
		} else {
			member.ID = id
			csvSlice = append(csvSlice, member)
			id++
		}
	}
	fmt.Println("completed CSV read: ", csvSlice)

	// Writing to a csv file
	file, err = os.Create("CrewMembers.csv")
	if err != nil {
		log.Fatal("Error creating crew file")
	}
	defer func() { _ = file.Close() }()
	// have to have a [][]string to write to a csv file, so add a to_csv method to the struct
	var csvWriter = csv.NewWriter(file)
	csvWriter.Comma = ';'
	if csvSlice == nil {
		log.Fatal("Empty crew list")
	}
	err = csvWriter.WriteAll(csvSlice.toCsv())
	if err != nil {
		log.Fatal("Error in writing crew list")
	}
}

