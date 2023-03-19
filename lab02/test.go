package main

import (
	"fmt"
)

func main() {
	fmt.Println(myFunction(1, 10, 3))
	test := 9.0 - 3.0/(1.0/3.0) + 1.0
	fmt.Println(test)
}

func zad3() {
	var a, b, c int
	fmt.Println("quadratic trinomial:")
	fmt.Scanln(&a, &b, &c)
	fmt.Println(a, b, c)
	del := (b * 2) - (4 * a * c)
	println(del)
}

func myFunction(a int, b int, c int) (result string) {
	delta := (b * b) - (4 * a * c)
	if delta >= 0 {
		s := string(delta)
		return s
	} else {
		return "false"
	}
}
