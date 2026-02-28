package main

import (
	"fmt"
	"time"
)

func writer(number int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := range number {
			ch <- i + 1
		}
		close(ch)
	}()

	return ch
}

func doubler(ch <-chan int) <-chan int {
	chNew := make(chan int)

	go func() {
		for i := range ch {
			chNew <- i * 2
			time.Sleep(500 * time.Millisecond)
		}
		close(chNew)
	}()

	return chNew
}

func readDoubleWrite() {

	ch1 := writer(10)
	ch2 := doubler(ch1)

	for i := range ch2 {
		fmt.Printf("Number: %d\n", i)
	}
}
