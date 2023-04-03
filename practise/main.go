package main

import (
	"fmt"
)

func main() {
	fmt.Println("practise")
	// noRepeating()
	onlyUnique()
}

func noRepeating() {
	var slice []int
	var sliceOut []int
	slice = append(slice, 1, 3, 3, 2, 3, 1)

	for i, el := range slice {
		if i == 0 {
			sliceOut = append(sliceOut, slice[0])
			continue
		}
		if el != sliceOut[len(sliceOut)-1] {
			sliceOut = append(sliceOut, slice[i])
		}
	}
	fmt.Println(slice)
	fmt.Println(sliceOut)
}

func onlyUnique() {
	slice := []int{1, 3, 3, 2, 3, 1}
	var sliceOut []int

	for i, el := range slice {
		exists := false
		if i == 0 {
			sliceOut = append(sliceOut, slice[0])
			continue
		}
		for _, el2 := range sliceOut {
			if el == el2 {
				exists = true
				break
			}
		}
		if !exists {
			sliceOut = append(sliceOut, slice[i])
		}

	}
	fmt.Println(slice)
	fmt.Println(sliceOut)
}
