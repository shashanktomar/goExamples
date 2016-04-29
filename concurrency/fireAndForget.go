package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fireAndForget(msg string) {
	go boring(msg)
	fmt.Println("Boring work starting...")
	time.Sleep(2 * time.Second)
	fmt.Println("I am leaving because you are boring.")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
