package main

import "fmt"

type CrewMember struct {
  name string
  age int
  address string
  rank string
  clearance int
}

func main() {
  kevin := CrewMember {
    name: "Kevin",
    age: 32,
    address: "here and now",
    rank: "sub-human",
    clearance: 1,
  }
  var jenny CrewMember
  jenny.name = "Jenny"
  jenny.age = 12
  jenny.address = "Mammas House"
  jenny.rank = "Boss of the staircase"
  jenny.clearance = 4

  var crew []CrewMember
  crew = append(crew, kevin, jenny, CrewMember{name: "Fondue", age: 1, address: "fridge", rank: "dinner", clearance: 0})

  for i, member := range crew {
    fmt.Println(i, ":", "name:", member.name, "  rank:", member.rank)
  }

  var crew_map map[string]CrewMember
  crew_map = make(map[string]CrewMember)
  crew_map = map[string]CrewMember{
    "kevin": kevin,
    "jenny": jenny,
  }
  crew_map["fondue"] = CrewMember{name: "Fondue", age: 1, address: "fridge", rank: "dinner", clearance: 0}
  fmt.Println("kevin:", crew_map["kevin"])
  delete(crew_member, "jenny")
  _, exists_bool := crew_map["jenny"]
  if !exists_bool { fmt.Println("should not exist in the map") }
}
