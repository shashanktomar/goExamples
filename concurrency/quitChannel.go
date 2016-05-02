package main

import (
	"fmt"
	"time"
)

func quitChannel() {
	quit := make(chan bool)
	go receiveCommunication(quit)
	time.Sleep(5 * time.Second)
	quit <- true
	fmt.Println("Let's Quit")
	// this is required to give some time to quit,
	// other program exit main function and async go routines never get a chance to read quit channel
	time.Sleep(500 * time.Millisecond)
}

func receiveCommunication(quit chan bool) {
	joe := someWork("joe ")
	ana := someWork("ana ")
	jeff := someWork("jeff ")
	for {
		select {
		case j := <-joe:
			fmt.Println(j)
		case a := <-ana:
			fmt.Println(a)
		case je := <-jeff:
			fmt.Println(je)
		case <-quit:
			fmt.Println("Quitting")
			return
		}
	}
}
