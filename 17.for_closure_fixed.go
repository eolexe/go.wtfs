package main

import (
	"fmt"
	"time"
)

func main() {
	input := []string{
		"lviv",
		"golang",
		"meetup",
	}

	for _, value := range input {
		//valueCopy := v
		go func(val string) {
			fmt.Println(val)
		}(value)
	}
	time.Sleep(1 * time.Second)
}

// The iteration variables in for statements are reused in each iteration.
// This means that each closure (aka function literal) created in
// your for loop will reference the same variable
// (and they'll get that variable's value at the
// time those goroutines start executing).
