package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectCase() {
	joe := someWork("joe ")
	ana := someWork("ana ")
	jeff := someWork("jeff ")
	dave := someWork("dave ")

	for {
		select {
		case j := <-joe:
			fmt.Println(j)
		case a := <-ana:
			fmt.Println(a)
		case je := <-jeff:
			fmt.Println(je)
		case d := <-dave:
			fmt.Println(d)
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("No communication received")
		}
	}
}

func someWork(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
	}()
	return c
}
