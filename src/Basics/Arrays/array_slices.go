package main

import "fmt"

const peopleCap int = 5

var listPeople [peopleCap]string
var slicePeople []string
var peopleLength int = 0

func addPersonToSlice(name ...string) {
	fmt.Println("Capacity of Slice List is ", cap(slicePeople))
	for _, i := range name {
		slicePeople = append(slicePeople, i)
	}
}

func getSliceListPersons() {
	fmt.Println("SliceList of People registered is as follows:")
	for _, i := range slicePeople {
		fmt.Println(i)
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
	addPersonToSlice("PersonA", "PersonB", "PersonC")
	addPersonToSlice("PersonD", "PersonE", "PersonF", "PersonG")
	getSliceListPersons()
}
