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
	v := banks["2"]
	fmt.Print(v[1])
	/*
		for i := range digtis {
			if digtis[i] == '0' || digtis[i] == '1' {
				log.Fatal("Error: A number less than 2 is present.")
			}
		}

		sli := make([]string, 0)
		var letterComb string
		digtisLen := len(digtis)

		for i := 0; i < digtisLen; i++ {

		}*/
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
