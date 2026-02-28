package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

func predictableTimeWork() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ch := make(chan bool)

	go func() {
		randomTimeWork()
		close(ch)

	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("function working longer than 3 sec")
	case <-ch:
		return nil
	}
}

func predictableTimeout() {
	err := predictableTimeWork()
	if err != nil {
		fmt.Println("predictableTimeout exited with error", err)
	} else {
		fmt.Println("predictableTimeout exited ok")

	}

}
