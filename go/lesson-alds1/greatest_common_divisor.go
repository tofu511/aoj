package main

import "fmt"

func main() {
	var x, y, greatestCommonDivisor int
	fmt.Scan(&x, &y)

	if x == y {
		greatestCommonDivisor = x
	} else if x >= y {
		divRemainder := x % y
		greatestCommonDivisor = findGreatestCommonDivisor(divRemainder, y)
	} else {
		divRemainder := y % x
		greatestCommonDivisor = findGreatestCommonDivisor(divRemainder, x)
	}

	fmt.Printf("%d\n", greatestCommonDivisor)
}

func findGreatestCommonDivisor(divRemainder int, smallerVal int) int {
	var result int
	for i := divRemainder; i > 0; i-- {
		if smallerVal % i == 0 && divRemainder % i == 0 {
			result = i
			break
		}
	}
	return result
}