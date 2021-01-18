package main

import (
	"FactoryPattern/dishes"
	"fmt"
)

func main() {
	fmt.Println("Enter preferred dish type")
	fmt.Println("0: Burger")
	fmt.Println("1: Noodles")
	fmt.Println("2: Pizza")

	var myType int
	fmt.Scan(&myType)

	var myDish, err = dishes.CreateDish(myType)

	if err == nil {
		myDish.LookIn()
		fmt.Println(myDish.GetContents())
	} else {
		fmt.Println(err)
	}
}
