package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Printf("a > b\n")
	} else if a < b {
		fmt.Printf("a < b\n")
	} else {
		fmt.Printf("a == b\n")
	}
}
