package main

import "fmt"

func main() {
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	row3 := []int{7, 8, 9}
	matrix := [][]int{row1, row2, row3}
	rev1 := []int{9, 8, 7}
	rev2 := []int{6, 5, 4}
	rev3 := []int{3, 2, 1}
	reverse := [][]int{rev1, rev2, rev3}

	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = reverse[i][j] + matrix[i][j]
		}
	}
	for i := range matrix {
		for _, k := range matrix[i] {
			fmt.Print(k)
			fmt.Print(" ")
		}
		println()
	}
}
