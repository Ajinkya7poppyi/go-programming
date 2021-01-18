/**********************************
binary search in go
**********************************/
package main

import "fmt"

func binarySearch(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}
	mid := int((lowIndex + highIndex) / 2)
	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binarySearch(array, target, mid, highIndex)
	} else {
		return mid
	}
}

func iterBinarySerach(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int
	for startIndex < endIndex {
		mid = int((startIndex + endIndex) / 2)
		if array[mid] > target {
			endIndex = mid
		} else if array[mid] < target {
			startIndex = mid
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	slice := []int{1, 2, 3, 5, 6, 10, 15, 18, 20, 25, 36, 45, 50}
	find := binarySearch(slice, 10, 0, len(slice)-1)
	fmt.Printf("Found the element by recursive binary search at index %d \n", find)
	find = iterBinarySerach(slice, 10, 0, len(slice)-1)
	fmt.Printf("Found the element by iterative binary search at index %d \n", find)
}
