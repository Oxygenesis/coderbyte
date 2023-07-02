package main

import (
	"fmt"
	"regexp"
	"strings"
)

func LongestWord(sen string) string {

	// code goes here
	expr := regexp.MustCompile(`[^\w\s]`)
	sen = expr.ReplaceAllString(sen, "")
	
	words := strings.Split(sen, " ")
	longest := ""
	var maxLen int

	for i := 0; i < len(words); i++ {
		if len(words[i]) > maxLen {
			maxLen = len(words[i])
			longest = words[i]
		}
	}

	return longest
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	testCase1 := "fun&!! time"
	fmt.Println(LongestWord(testCase1))

}
