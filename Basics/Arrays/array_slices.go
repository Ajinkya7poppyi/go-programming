package main

import "fmt"

const peopleCap int = 5

var listPeople [peopleCap]string
var slicePeople []string
var peopleLength int = 0

func addPersonToSlice(name string) {
	fmt.Println("Capacity of Slice List is ", cap(slicePeople))
	slicePeople = append(slicePeople, name)
}

func getSliceListPersons() {
	fmt.Println("SliceList of People registered is as follows:")
	for i := 0; i < len(slicePeople); i++ {
		fmt.Println(slicePeople[i])
	}
}

func addPerson(name string) {
	if peopleLength < peopleCap {
		listPeople[peopleLength] = name
		peopleLength++
	} else {
		fmt.Println("Too Many Items. We are closed now!")
	}
}

func getListPersons() {
	fmt.Println("List of people registered is as follows:")
	for i := 0; i < peopleLength; i++ {
		fmt.Println(listPeople[i])
	}
}

func main() {
	addPerson("PersonA")
	addPerson("PersonB")
	addPerson("PersonC")
	addPerson("PersonD")
	addPerson("PersonE")
	addPerson("PersonF")
	addPerson("PersonF")
	getListPersons()
	addPersonToSlice("PersonA")
	addPersonToSlice("PersonB")
	addPersonToSlice("PersonC")
	addPersonToSlice("PersonD")
	addPersonToSlice("PersonE")
	addPersonToSlice("PersonF")
	addPersonToSlice("PersonG")
	getSliceListPersons()
}
