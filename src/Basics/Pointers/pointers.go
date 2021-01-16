package main

import "fmt"

// ChangeX   Change value of X
func ChangeX(X *int) {
	*X = 20
}

func main() {
	var X int
	X = 6
	ChangeX(&X)
	fmt.Println("The Value of X is:", X)
}
