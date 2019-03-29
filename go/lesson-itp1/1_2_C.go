package main

import (
	"fmt"
	"sort"
)

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	arr := []int{x, y, z}
	sort.Ints(arr)
	fmt.Printf("%d %d %d\n", arr[0], arr[1], arr[2])
}
