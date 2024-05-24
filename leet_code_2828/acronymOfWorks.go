package main

import "fmt"

func main() {
	var words = make([]string, 0)
	var acronym = "h"

	words = append(words, "hello", "world")

	fmt.Println(findAcronym(words, acronym))
}

func findAcronym(w []string, acronymLetter string) bool {
	if len(w) != len(acronymLetter) {
		return false
	}

	for i, word := range w {
		if word[0] != acronymLetter[i] {
			return false
		}
	}

	return true
}
