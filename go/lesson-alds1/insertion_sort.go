package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	inputSlice := stdinToIntSlice()

	insertionSort(inputSlice, n)
}

func insertionSort(inputSlice []int, n int)  {
	for i := 1; i <= n - 1; i++ {
		printSlice(inputSlice)
		v := inputSlice[i]
		j := i - 1

		for j >= 0 && inputSlice[j] > v {
			inputSlice[j + 1] = inputSlice[j]
			j--
		}
		inputSlice[j + 1] = v
	}
	printSlice(inputSlice)
}

func stdinToIntSlice() []int  {
	stdin := bufio.NewScanner(os.Stdin)
	var res []int
	stdin.Scan()
	text := stdin.Text()
	texts := strings.Split(text, " ")
	for _, t := range texts {
		intValue, _ := strconv.Atoi(t)
		res = append(res, intValue)
	}

	return res
}

func printSlice(s []int)  {
	for i, val := range s {
		if i + 1 != len(s) {
			fmt.Printf("%d ", val)
		} else {
			fmt.Printf("%d", val)
		}
	}
	fmt.Printf("\n")
}