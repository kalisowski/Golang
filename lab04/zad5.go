package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var a = [][]int{{1, 2}, {3, 4}, {5, 6}}
	var b = [][]int{{1, 2, 3}, {4, 5, 6}}
	var c, err = matrixMult(a, b)
	var d, err2 = matrixMult(b, a)
	if err != nil || err2 != nil {
		log.Fatalln(err, err2)
	}
	for i := range c {
		fmt.Println(c[i])
	}
	fmt.Println()
	for i := range d {
		fmt.Println(d[i])
	}
}

func matrixMult(a, b [][]int) ([][]int, error) {
	if len(a) == 0 || len(b) == 0 || len(a[0]) != len(b) {
		return nil, errors.New("invalid matrix sizes")
	}
	var c = make([][]int, len(a))
	for i := range c {
		c[i] = make([]int, len(b[0]))
		for j := range c[i] {
			for k := range a[i] {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c, nil
}
