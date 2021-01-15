package main

import "fmt"

func main() {
	age := 69
	if age < 13 {
		fmt.Println("Wow You're Young!")
	} else if age < 20 {
		fmt.Println("You're Teenager")
	} else if age < 30 {
		fmt.Println("You're in your Twenties")
	} else if age < 40 {
		fmt.Println("You're in yure Thrities")
	} else if age < 50 {
		fmt.Println("You're Getting There")
	} else {
		fmt.Println("You're Over The Hill")
	}
}
