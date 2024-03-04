// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
package main

import (
	"fmt"
)

func main() {
	var digits string = "23"
	filter(digits)
}

func filter(digtis string) {
	banks := whichBanks(digtis)
	fmt.Println(banks)
	var letterComb []string

	// We will make our combinations, out of the first element in our map(hash table).
	firstMapElementLen := len(banks[string(digtis[0])])
	firstMapElement := banks[string(digtis[0])]
	for i := 0; i < firstMapElementLen; i++ {
		tempt1 := firstMapElement[i]
		fmt.Println("Element 1:", tempt1)
		// ToDo: Maybe it would be a good idea to make the below code a function.
		for j := 1; j < len(banks); j++ { // Plus 1 for the length, so that we don't interact with the first map element.
			var tempt2 string = tempt1 + banks[string(digtis[j])][j-1]
			fmt.Println("Element 2:", tempt2)
			letterComb = append(letterComb, tempt2)
		}
	}

	fmt.Println(letterComb)

}

func whichBanks(digtis string) map[string][]string {
	banks := make(map[string][]string)

	// Arrays with letters corresponding to each digit.
	bank2 := [3]string{"a", "b", "c"}
	bank3 := [3]string{"d", "e", "f"}
	bank4 := [3]string{"g", "h", "i"}
	bank5 := [3]string{"j", "k", "l"}
	bank6 := [3]string{"m", "n", "o"}
	bank7 := [4]string{"p", "q", "r", "s"}
	bank8 := [3]string{"t", "u", "v"}
	bank9 := [4]string{"w", "x", "y", "z"}

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
