package main

import (
	"fmt"
)

type test struct {
	pole int
}

type test2 struct {
	cztery []int
}

type razz struct {
	dwa RuneMap
}

type Map map[string]int
type RuneMap map[rune]test2
type BoolMap map[bool]test2

func main() {
	fmt.Println("Hello, 🦆")
	raz := razz{RuneMap{'3': test2{[]int{1, 2, 3}}}}
	fmt.Println(raz.dwa['3'].cztery[0])
	trzy := BoolMap{false: test2{[]int{1, 2, 3}}}
	fmt.Println(trzy[false].cztery[0])
	mapa := Map{"a": 1, "b": 2}
	dane1 := []Map{mapa, mapa, mapa}
	fmt.Println(dane1[0]["a"])
	var dane [2][2]test
	dane[0][0].pole = 1
	fmt.Println(dane[0][0].pole)

	fmt.Println("practise")
	// noRepeating()
	// onlyUnique()
	test := [...]int{1, 2, 3, 4, 5}
	// test2 := [...]int{1, 2, 3, 4, 5}Int

	fmt.Println(len(test))
	fmt.Println(Int(1))
	fmt.Println(Int(1))
}

func Int(i int) (a int) {
	a = i
	return
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
