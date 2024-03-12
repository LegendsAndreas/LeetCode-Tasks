package main

import (
	"fmt"
	"log"
)

func main() {
	i := 2.0
	j := -2147483648

	fmt.Println(myPow(i, j))
}

func myPow(x float64, n int) float64 {
	if x == 0 && n < 0 {
		log.Fatal("'x' is 0 and 'n' is less than 0")
	}

	// Anything to the power of 0 is just 1.
	if n == 0 {
		fmt.Println("n is 0!")
		return 1.0
	}

	// 1 to the power of anything, is always 1.
	if x == 1 {
		fmt.Println("x is 1!")
		return 1
	}

	var powered float64 = 1
	if n > 1 {
		powered = powerFunc(x, n)
	} else {
		n = -n
		powered = powered / (powerFunc(x, n))
	}

	return powered
}

func powerFunc(x float64, n int) float64 {
	var powder float64 = 1
	for i := 0; i < n; i++ {
		powder = powder * x
	}
	return powder
}
