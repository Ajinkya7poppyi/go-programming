package main

import "fmt"

func printAdd(i int, j int) {
	fmt.Println("Addition of ", i, " and ", j, " is ", i+j)
}

func returnAdd(i int, j int) int {
	return i + j
}

func add(i int, j int, k *int) {
	*k = i + j
}
func main() {
	printAdd(20, 30)
	fmt.Println("Addition from returnadd function is ", returnAdd(2, 85))
	var k int
	add(54, 89, &k)
	fmt.Println("Addition by passing pointer is ", k)
}
