package main

import "fmt"

func main() {
	m := map[int]string{
		1: "January",
		2: "February",
		3: "March",
		4: "April",
		5: "May",
		6: "June",
	}

	for key, value := range m {
		fmt.Println(key, value)
	}
}
