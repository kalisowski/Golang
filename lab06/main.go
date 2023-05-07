package main

import (
	"flag"
	"fmt"
	"math"
)

type t struct{ a, b, c float64 }
type w struct {
	x [2]float64
	d bool
}

func main() {
	A, B, C := flag.Float64("a", 0, ""), flag.Float64("b", 0, ""), flag.Float64("c", 0, "")
	flag.Parse()
	T := t{*A, *B, *C}
	W := calc(&T)
	if W.d {
		fmt.Println(W.x[0], W.x[1])
	} else {
		fmt.Println("!")
	}
}

func calc(T *t) w {
	S := T.b*T.b - 4*T.a*T.c
	if T.a == 0 {
		return w{[2]float64{0, 0}, false}
	}
	if S < 0 {
		return w{[2]float64{0, 0}, false}
	}
	V := math.Sqrt(S)
	return w{[2]float64{(-T.b - V) / (2 * T.a), (-T.b + V) / (2 * T.a)}, true}
}
