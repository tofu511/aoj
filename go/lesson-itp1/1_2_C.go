package main

import (
	"fmt"
	"sort"
)

// https://onlinejudge.u-aizu.ac.jp/courses/lesson/2/ITP1/2/ITP1_2_C
func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	arr := []int{x, y, z}
	sort.Ints(arr)
	fmt.Printf("%d %d %d\n", arr[0], arr[1], arr[2])
}
