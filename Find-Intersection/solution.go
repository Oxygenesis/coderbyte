package main

import (
	"fmt"
	"strings"
)

func FindIntersection(strArr []string) string {

	// code goes here
	var (
		res  string
		arrA = stringToArr(strArr[0])
		arrB = stringToArr(strArr[1])
		mapB = make(map[string]bool, len(arrB))
	)

	if len(strArr[0]) == 0 || len(strArr[1]) == 0 {
		return "false"
	}

	for i := 0; i < len(arrB); i++ {
		mapB[arrB[i]] = true
	}

	for i := 0; i < len(arrA); i++ {
		if mapB[arrA[i]] {
			res += arrA[i] + ","
		}
	}
	res = strings.TrimSuffix(res, ",")
	if res == "" {
		res = "false"
	}
	return res
}

func stringToArr(str string) []string {
	return strings.Split(strings.Replace(str, " ", "", -1), ",")
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	input := []string{"1, 3, 4, 7, 13", "1, 2, 4, 13, 15"}
	fmt.Println(FindIntersection(input))
}
