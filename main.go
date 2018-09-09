package main

import "fmt"

func main() {
	arr := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	target := 33

	fmt.Println(BinarySearch(arr, target))
}

func BinarySearch(array []int, target int) int {
	return BinarySearchHelper(array, target, 0, len(array)-1)
}

func BinarySearchHelper(array []int, target int, left int, right int) int {
	if left > right {
		// target int not found
		return -1
	}

	// mid is a pivot
	mid := (left + right) / 2

	potentialMatch := array[mid]

	if target == potentialMatch {
		return mid
	}

	if target < potentialMatch {
		return BinarySearchHelper(array, target, left, mid-1)
	}

	return BinarySearchHelper(array, target, mid+1, right)

}
