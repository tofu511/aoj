package main

import "fmt"

// https://onlinejudge.u-aizu.ac.jp/courses/lesson/2/ITP1/2/ITP1_2_A
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
