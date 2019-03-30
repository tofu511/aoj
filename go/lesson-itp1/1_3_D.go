package main

import (
	"fmt"
)

// https://onlinejudge.u-aizu.ac.jp/courses/lesson/2/ITP1/all/ITP1_3_D
func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	count := 0

	for x <= y {
		if z % x == 0 {
			count++
		}
		x++
	}
	fmt.Printf("%d\n", count)
}
