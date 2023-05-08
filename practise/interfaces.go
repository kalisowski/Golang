package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, 🦆")
	newSlice := MySlice{Int(1), String("2"), Int(3), String("4")}
	fmt.Println(newSlice)

}

type IntOrString interface {
	isIntOrString()
}

type Int int

func (i Int) isIntOrString() {}

type String string

func (s String) isIntOrString() {}

type MySlice []IntOrString

type Addable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

type AddableVal[T Addable] struct {
	X T
}

func (p AddableVal[T]) Add(q AddableVal[T]) AddableVal[T] {
	return AddableVal[T]{p.X + q.X}
}

type Addable2 interface {
	int | float64
}
