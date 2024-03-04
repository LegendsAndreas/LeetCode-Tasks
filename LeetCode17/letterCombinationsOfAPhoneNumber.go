// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
package main

import (
	"fmt"
)

func main() {
	var digits string = "234"
	letterCombs := filter(digits)
	fmt.Println(letterCombs)
}

func filter(digtis string) []string {
	banks := whichBanks(digtis)
	var letterComb []string

	// In case the length of banks is less than 2, we can just return it right away.
	if len(banks) < 2 {
		return banks[digtis]
	}

	// We will make our combinations, out of the first element in our map(hash table), by addressing the first digit.
	firstMapElementLen := len(banks[string(digtis[0])])
	firstMapElement := banks[string(digtis[0])] // We set firstMapElement to be equal to the first array in our map.

	for i := 0; i < firstMapElementLen; i++ {
		temptLetter := firstMapElement[i]
		for j := 1; j < len(banks); j++ { // Plus 1 to not hit the first array in our map.
			insertLetterCombs(&letterComb, banks[string(digtis[j])], temptLetter)
		}
	}

	return letterComb

}

// whichBanks is a function that takes in a string (digits) and returns a map of strings to slices of strings (banks).
// It populates the map based on the digits provided and assigns the corresponding letter banks to each digit.
// It uses a switch statement to determine the digit and assigns the letter bank accordingly.
// It then returns the populated map.
func whichBanks(digtis string) map[string][]string {
	banks := make(map[string][]string) // The variable which will return the relevant banks.

	// Arrays with letters corresponding to each digit.
	bank2 := [3]string{"a", "b", "c"}
	bank3 := [3]string{"d", "e", "f"}
	bank4 := [3]string{"g", "h", "i"}
	bank5 := [3]string{"j", "k", "l"}
	bank6 := [3]string{"m", "n", "o"}
	bank7 := [4]string{"p", "q", "r", "s"}
	bank8 := [3]string{"t", "u", "v"}
	bank9 := [4]string{"w", "x", "y", "z"}

	// Insert relevant banks into the bank variable.
	for i := 0; i < len(digtis); i++ {
		switch digtis[i] {
		case '2':
			banks["2"] = bank2[:]
		case '3':
			banks["3"] = bank3[:]
		case '4':
			banks["4"] = bank4[:]
		case '5':
			banks["5"] = bank5[:]
		case '6':
			banks["6"] = bank6[:]
		case '7':
			banks["7"] = bank7[:]
		case '8':
			banks["8"] = bank8[:]
		case '9':
			banks["9"] = bank9[:]

		}
	}

	return banks
}

// insertLetterCombs is a function that takes in a pointer to a slice of strings (letterCombs), a slice of strings (currentBank), and a string (currentMainLetter). It appends combinations.
// I thought it felt better by not returning anything and just using a pointer.
func insertLetterCombs(letterCombs *[]string, currentBank []string, currentMainLetter string) {
	for i := 0; i < len(currentBank); i++ {
		tempt := currentMainLetter + currentBank[i]
		*letterCombs = append(*letterCombs, tempt)
	}
}
