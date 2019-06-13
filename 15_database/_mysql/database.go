package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // native fast mysql driver written in golang, init function needs to run
	"log"
	"strings"
)

type CrewMember struct {
	ID 				  	int		`mysql:"id"`
	Name              	string	`mysql:"name"`
	SecurityClearance 	int
	Position 	  		string
}
type Crew []CrewMember

func main()  {
	db, err :=sql.Open("mysql", "morb:ch1ck3n@/hydra?parseTime=true") // translate times to go time objects
	if err != nil {
		log.Fatal("Oh dear our DB was not connectable.")
	}
	defer func(){ _ = db.Close() }()

	var captain = getCrewByPosition(db, []string{"'Captain'"})[0]
	var positions = []string{"'Frogman'","'Chef'"} // have to surround strings with ' so can join correctly into sql string list
	var crew = getCrewByPosition(db, positions)

	fmt.Printf("hello %+v\n", captain)
	fmt.Println("here is your selection:")
	fmt.Println("-----------------------")
	 for _, member := range crew {
		 fmt.Printf("available for hire: %+v\n", member)
	 }
	 fmt.Println("Added crew number: ", addCrewMember(db, CrewMember{Name: "Anguel McFarminster", SecurityClearance:6, Position:"Digger"}))
}

func getCrewByPosition(db *sql.DB, positions []string) (crew Crew) {
	var query = fmt.Sprintf(`select * from personnel where Position in (%s)`, strings.Join(positions, ","))
	rows, _ := db.Query(query)
	for rows.Next() {
		var member CrewMember
		err := rows.Scan(&member.ID, &member.Name, &member.SecurityClearance, &member.Position)
		if err != nil {
			log.Fatal("Could not find a frogman or chef in the DB", err)
		}
		crew = append(crew, member)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func addCrewMember(db *sql.DB, member CrewMember) int64 {
	result, err := db.Exec(
		"insert into personnel (name, SecurityClearance, Position) values (?,?,?)",
		member.Name, member.SecurityClearance, member.Position)
	if err != nil {
		log.Fatal("Could not insert: ", member)
	}
	affected, _ := result.RowsAffected()
	id, _ := result.LastInsertId()
	log.Println("Added", affected, "new rows")
	return id
}

//func addCrewMembers(db *sql.DB, crew Crew) {
//	transaction, err := db.Begin()
//	if err != nil {
//		log.Fatal("Could not begin transaction")
//	}
//	statement, err := transaction.Prepare("insert into personnel (name, SecurityClearance, Position) values (?,?,?)")
//	if err != nil {
//		_ = transaction.Rollback()
//		log.Fatal("Could not prepare statement")
//	}
//
//
//}