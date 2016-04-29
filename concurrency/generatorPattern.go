package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Here a function which run in goroutine return you the channel on which it will communicate

func generator() {
	joe := boringGenerator("joe ")
	ana := boringGenerator("ana ")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-joe)
		fmt.Printf("You say: %q\n", <-ana)
	}
}

func boringGenerator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
