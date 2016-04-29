package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// fireAndForget("I am boring: ")
	// streamingChannel("Yippy channels: ")
	// generator()
	// multiplexer()
	// seqMultiplexer()
	// selectCase()
	selectTimeout()
}
