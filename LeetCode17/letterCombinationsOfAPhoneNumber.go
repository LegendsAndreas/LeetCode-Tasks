package main

import "log"

func main() {
	var digits string = "23"
	filter(digits)

}

func filter(digtis string) {
	// Arrays with the corrosponding

	bank2 := [3]string{"a", "b", "c"}
	bank3 := [3]string{"d", "e", "f"}
	bank4 := [3]string{"g", "h", "i"}
	bank5 := [3]string{"j", "k", "l"}
	bank6 := [3]string{"m", "n", "o"}
	bank7 := [4]string{"p", "q", "r", "s"}
	bank8 := [3]string{"t", "u", "v"}
	bank9 := [4]string{"w", "x", "y", "z"}

	//enteredDigitsLen := len(enteredDigits)
	for i := range digtis {
		if digtis[i] == '0' || digtis[i] == '1' {
			log.Fatal("Error: A number less than 2 is present.")
		}
	}

	sli := make([]string,0)
	var letterComb string

	for i := 0; i < ; i++ {
		
	}
	
}
