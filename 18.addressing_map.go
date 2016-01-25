package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	people := map[string]Person{
		"mike": {"Michael"},
	}

	fmt.Println(people["mike"])
	people["mike"].Name = "Mikey"
}
