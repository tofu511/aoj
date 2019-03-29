package main

import "fmt"

// https://onlinejudge.u-aizu.ac.jp/courses/lesson/2/ITP1/2/ITP1_2_B
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a < b && b < c {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}