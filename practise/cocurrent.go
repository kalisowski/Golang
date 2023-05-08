package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, playground")

	testchannel := make(chan int, 2)
	testchannel <- 21
	testchannel <- 37

	go func() {
		testchannel <- 1
		testchannel <- 2
		testchannel <- 3
		close(testchannel)
	}()

	go func() {
		fmt.Println("test")
	}()
	time.Sleep(1 * time.Second)
}
