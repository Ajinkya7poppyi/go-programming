package main

import "fmt"

// Set sets
type Set map[string]struct{}

func getSetValues(s Set) []string {
	var retVal []string
	for k := range s {
		retVal = append(retVal, k)
	}
	return retVal
}

func main() {
	s := make(Set)
	//Add items
	s["item1"] = struct{}{}
	s["item2"] = struct{}{}
	//get and print items
	fmt.Println(getSetValues(s))

}
