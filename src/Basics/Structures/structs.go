package main

import (
	"fmt"
)

//SMember Structure
type SMember struct {
	name       string
	age        int
	address    string
	mobile     string
	rollnumber int
}

func main() {
	// Declare a struct
	sm1 := SMember{"Ajinkya", 29, "Pune", "9028077960", 23}

	sm2 := SMember{
		name:       "Arjun",
		age:        22,
		address:    "Mumbai",
		mobile:     "569556565",
		rollnumber: 25,
	}

	var sm3 SMember
	sm3.name = "Ashitosh"
	sm3.age = 29
	sm3.mobile = "44545455454"
	sm3.address = "Pune"
	sm3.rollnumber = 89

	var classroom []SMember
	classroom = append(classroom, sm1, sm2, sm3)

	for i, v := range classroom {
		fmt.Println(i+1, v)
	}
}
