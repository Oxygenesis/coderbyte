package main

import "fmt"

func BracketMatcher(str string) string {

	// code goes here
	res := "0"
	mapBracket := make(map[string]int, 1)
	openBracket := "("
	for _, v := range str {
		char := string(v)
		if char == openBracket {
			mapBracket[openBracket]++
			continue
		}

		if char == ")" {
			if mapBracket[openBracket] <= 0 {
				return res
			}
			mapBracket[openBracket]--
		}
	}

	if mapBracket[openBracket] == 0 {
		res = "1"
	}
	return res
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(BracketMatcher("hello()"))

}
