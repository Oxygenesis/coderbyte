package main

import "fmt"

func main() {
	// Test the function with sample inputs
	arr1 := []int{5, 10, 15, 20, 25}         // Arithmetic sequence
	arr2 := []int{2, 4, 8, 16, 32}           // Geometric sequence
	arr3 := []int{1, 2, 3, 4, 5, 10, 20, 30} // Not arithmetic or geometric

	fmt.Println(ArithGeo(arr1)) // Output: Arithmetic
	fmt.Println(ArithGeo(arr2)) // Output: Geometric
	fmt.Println(ArithGeo(arr3)) // Output: -1
	testCases := [][]int{
		{1, 3, 5, 7, 9},                 // Expected: "Arithmetic"
		{2, 4, 8, 16, 32, 64},           // Expected: "Geometric"
		{1, 2, 5, 10, 20},               // Expected: "-1"
		{0, 0, 0, 0, 0},                 // Expected: "Arithmetic"
		{1, 1, 1, 1, 1, 1},              // Expected: "Geometric"
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, // Expected: "Arithmetic"
	}

	for _, testCase := range testCases {
		result := ArithGeo(testCase)
		fmt.Println(result)
	}
}

func ArithGeo(arr []int) string {
	// Check if the array is arithmetic or geometric
	if len(arr) < 2 {
		return "-1"
	}
	if isArithmetic(arr) {
		return "Arithmetic"
	} else if isGeometric(arr) {
		return "Geometric"
	} else {
		return "-1"
	}
}

func isArithmetic(arr []int) bool {
	diff := arr[1] - arr[0]
	for i := 1; i < len(arr)-2; i++ {
		if arr[i+1]-arr[i] != diff {
			return false
		}
	}
	return true
}

func isGeometric(arr []int) bool {
	diff := arr[1] / arr[0]
	for i := 1; i < len(arr)-2; i++ {
		if arr[i+1] != arr[i]*diff {
			return false
		}
	}
	return true
}
