package main

import "fmt"

// https://onlinejudge.u-aizu.ac.jp/courses/lesson/2/ITP1/2/ITP1_2_D
func main() {
	var w, h, x, y, r int
	fmt.Scan(&w, &h, &x, &y, &r)

	xAxis := x - r >= 0 && x + r <= w
	yAxis := y - r >= 0 && y + r <= h

	if xAxis && yAxis {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}
