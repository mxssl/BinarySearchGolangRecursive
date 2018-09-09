package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	arr := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	target := 33
	expected := 3

	output := BinarySearch(arr, target)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}

	fmt.Printf("Expected: %v, Output: %v\n", expected, output)
}

func TestCase2(t *testing.T) {
	arr := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	target := 75
	expected := -1

	output := BinarySearch(arr, target)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}

	fmt.Printf("Expected: %v, Output: %v\n", expected, output)
}
