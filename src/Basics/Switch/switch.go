package main

import "fmt"

func main() {
	for num := 0; num <= 11; num++ {
		fmt.Println("Number is", num)
		switch num {
		case 0:
			fmt.Println("This is Number ZERO")
		case 1:
			fmt.Println("This is Number ONE")
		case 2:
			fmt.Println("This is Number TWO")
		case 3:
			fmt.Println("This is Number THREE")
		case 4:
			fmt.Println("This is Number FOUR")
		case 5:
			fmt.Println("This is Number FIVE")
		case 6:
			fmt.Println("This is Number SIX")
		case 7:
			fmt.Println("This is Number SEVEN")
		case 8:
			fmt.Println("This is Number EIGHT")
		case 9:
			fmt.Println("This is Number NINE")
		case 10:
			fmt.Println("This is Number TEN")
		default:
			fmt.Println("Number Not In List")
		}
	}
}
