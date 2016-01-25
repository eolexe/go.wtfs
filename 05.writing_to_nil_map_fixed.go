package main

import "fmt"

// Why?
// Map types are reference types, like pointers or slices,
// and so the value of m above is nil;
// it doesn't point to an initialized map.

func main() {
	//m := make(map[string]int)
	m := map[string]int{}
	m["foo"] = 1
	fmt.Println(m)
}
