package main

import (
	"fmt"
	"math"
)

func main() {
	x := 88.0123456789
	result := round(x)
	fmt.Println(result)
}

func round(input float64) (wholeNumber float64) {
	var fractionalNumber float64
	if input != 0 {
		wholeNumber, fractionalNumber = math.Modf(input)
		fmt.Println("round:", wholeNumber, fractionalNumber)
		return
	}
	return 0
}
