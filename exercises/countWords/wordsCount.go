package main

import (
	"fmt"
	"strings"
)

// Counts the number of occurrences of each word in a given string.
// It returns a map where the keys are the words and the values are the counts.
func WordsCount(s string) map[string]int {

	words := strings.Fields(s)
	result := make(map[string]int)
	for i := 0; i < len(words); i++ {
		result[words[i]] = result[words[i]] + 1
	}
	return result
}

func main() {
	fmt.Println(WordsCount("My name is Rodrigo and I am learning Go!"))
	fmt.Println(WordsCount("Go Js Js Java Java Java C# C# C#"))
	fmt.Println(WordsCount(""))
}
