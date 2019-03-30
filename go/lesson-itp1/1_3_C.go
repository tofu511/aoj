package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://onlinejudge.u-aizu.ac.jp/courses/lesson/2/ITP1/all/ITP1_3_C
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		inputs := strings.Split(stdin.Text(), " ")
		x, _ := strconv.Atoi(inputs[0])
		y, _ := strconv.Atoi(inputs[1])

		if x == 0 && y == 0 {
			break
		}

		if x < y {
			fmt.Printf("%d %d\n", x, y)
		} else {
			fmt.Printf("%d %d\n", y, x)
		}
	}
}
