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
		go func() {
			fmt.Println(value)
		}()
		//time.Sleep(1 * time.Second)
	}
	time.Sleep(1 * time.Second)
}
