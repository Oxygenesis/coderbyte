package main

import (
	"fmt"
	"sort"
)

func main() {
	// Test the Consecutive function
	arr := []int{4, 8, 6}
	result := Consecutive(arr)
	fmt.Println(result) // Expected output: 2

	arr = []int{5, 10, 15}
	result = Consecutive(arr)
	fmt.Println(result) // Expected output: 8

	testCases := [][]int{
		{4, 8, 6},            // Expected: 2
		{5, 10, 15},          // Expected: 8
		{-3, -2, 0, 1},       // Expected: 3
		{100, 101, 102, 105}, // Expected: 2
		{-10, -8, -6, -5},    // Expected: 4
	}

	for _, testCase := range testCases {
		result := Consecutive(testCase)
		fmt.Println(result)
	}
}

func Consecutive(arr []int) int {
	sort.Ints(arr)
	min, max := arr[0], arr[len(arr)-1]
	rangeData := max - (min + 1)
	extraData := len(arr) - 2 //min and max
	return rangeData - extraData
}
