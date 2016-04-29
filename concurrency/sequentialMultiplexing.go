package main

import (
	"fmt"
	"math/rand"
	"time"
)

type message struct {
	msg  string
	wait chan bool
}

func seqMultiplexer() {
	joe := waitingGenerator("joe ")
	ana := waitingGenerator("ana ")
	club := fanInMsg(joe, ana)
	for i := 0; i < 10; i++ {
		m1 := <-club
		fmt.Printf("You say: %q\n", m1.msg)
		m2 := <-club
		fmt.Printf("You say: %q\n", m2.msg)
		m1.wait <- true
		m2.wait <- true
	}
}

func waitingGenerator(msg string) <-chan message {
	mc := make(chan message)
	wait := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			mc <- message{fmt.Sprintf("%s %d", msg, i), wait}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-wait
		}
	}()
	return mc
}

func fanInMsg(channnels ...<-chan message) <-chan message {
	ch := make(chan message)
	for i := range channnels {
		c := channnels[i]
		go func() {
			for {
				ch <- <-c
			}
		}()
	}
	return ch
}
