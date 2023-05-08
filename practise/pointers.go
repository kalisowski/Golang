package main

import "fmt"

func main() {
	var y int = 15
	var pointer *int = &y

	fmt.Println(pointer)  // prints the memory address of y
	fmt.Println(*pointer) // prints the value of y
	fmt.Println(&pointer) // prints the memory address of pointer

	t := &y        // &y is the memory address of y
	fmt.Println(t) // prints the memory address of y
	r := *t        // *t is the value of y
	fmt.Println(r) // prints 15 (the value of y)

	z := *&y       // *&y is equivalent to y
	fmt.Println(z) // prints 15 (the value of y)

	b := &*&y // &*&y is equivalent to &y
	fmt.Println(b)
	c := *&*&y // *&*&y is equivalent to *&y
	fmt.Println(c)

	var x int = 10
	var ptr *int = &x
	fmt.Println(*ptr) // prints the value of x
	fmt.Println(ptr)  // prints the memory address of x

	// Access memory address and dereference it in a single line
	fmt.Println(*(&x))   // prints the value of x
	fmt.Println(*(&ptr)) // prints the memory address of x
}
