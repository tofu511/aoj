package main

import (
	"fmt"
)

// https://onlinejudge.u-aizu.ac.jp/courses/lesson/2/ITP1/1/ITP1_1_C
func main() {
	var length, width int
	fmt.Scan(&length, &width)
	fmt.Printf("%d %d\n", length * width, 2 * (length + width))
}
