package main

import "fmt"

func main() {
	months := map[string]struct{}{
		"January":  struct{}{},
		"February": struct{}{},
		"March":    struct{}{},
		"April":    struct{}{},
		"May":      struct{}{},
		"June":     struct{}{},
	}

	if _, ok := months["March"]; ok {
		fmt.Println("Found!")
	}
}
