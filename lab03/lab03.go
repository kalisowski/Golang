package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	res := collatz(100)
	fmt.Println(res)
	println(collatzLen(10, 100))
	println(collatzLen(1000, 2000))
	println(collatzLen(2000, 3000))
	println(collatzLen(3000, 4000))
	println(collatzLen(30000, 40000))
	f, err := os.Create("collatz.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for i := 1; i <= 100000; i++ {
		length := len(collatz(i))
		fmt.Fprintf(f, "%d %d\n", i, length)
	}
}

func collatz(n int) []int {
	if n == 1 {
		return []int{1}
	}
	if n%2 == 0 {
		return append(collatz(n/2), n)
	}
	result := append(collatz(3*n+1), n)
	return result
}

func collatzLen(start int, end int) (int, int) {
	var maxLen, maxValue int
	for i := start; i <= end; i++ {
		list := collatz(i)
		if len(list) > maxLen {
			maxLen = len(list)
			maxValue = i
		}
	}
	return maxValue, maxLen
}
