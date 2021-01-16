package main

import "fmt"

// ChangeX   Change value of X
func ChangeX(X *int) {
	*X = 20
}

func main() {
	var X int
	X = 6
	fmt.Println("The Value of X Before is:", X)
	ChangeX(&X)
	fmt.Println("The Value of X After is:", X)
}
