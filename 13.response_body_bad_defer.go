package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com2")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error.")
		return
	}
	fmt.Println(resp.StatusCode)
}
