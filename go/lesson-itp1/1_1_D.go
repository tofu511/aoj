package main

import "fmt"

// https://onlinejudge.u-aizu.ac.jp/courses/lesson/2/ITP1/1/ITP1_1_D
func main() {
	var x int
	fmt.Scan(&x)
	h := x / 60 / 60
	m := (x / 60) % 60
	s := x % 60
	fmt.Printf("%d:%d:%d\n", h, m, s)
}
