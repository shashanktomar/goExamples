package main

import (
	"fmt"
	"math/rand"
	"time"
)

func streamingChannel(msg string) {
	c := make(chan string)
	go boringWithChannel(msg, c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("I am leaving because you are boring.")
}

func boringWithChannel(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
