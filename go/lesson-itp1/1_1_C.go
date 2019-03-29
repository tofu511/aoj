package main

import (
	"fmt"
)

func main() {
	var length, width int
	fmt.Scan(&length, &width)
	fmt.Printf("%d %d\n", length * width, 2 * (length + width))
}
