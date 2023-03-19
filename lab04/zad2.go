package main

import "fmt"

func main() {
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	row3 := []int{7, 8, 9}
	matrix := [][]int{row1, row2, row3}

	for i := range matrix {
		for _, k := range matrix[i] {
			fmt.Print(k)
			fmt.Print(" ")
		}
		println()
	}
}
