package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://onlinejudge.u-aizu.ac.jp/courses/lesson/2/ITP1/all/ITP1_3_B
func main() {
	stdin := bufio.NewScanner(os.Stdin)
	count := 1
	for stdin.Scan() {
		inputText := stdin.Text()
		inputInt, _ := strconv.Atoi(inputText)
		if inputInt == 0 {
			break
		}
		fmt.Printf("Case %d: %d\n", count, inputInt)
		count++
	}
}
