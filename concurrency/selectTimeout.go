package main

import (
	"fmt"
	"time"
)

// Here timer is created every time the loop starts waiting for communication. If any other communication
// takes more time than timer, the timer's case will run which return from the function
func selectTimeoutOnEachWait() {
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
		case <-time.After(1 * time.Second):
			fmt.Println("I am leaving, you are slow")
			return
		}
	}
}

// Here timer is created outside, the loop will return after timer communicate.
func selectOverallTimeout() {
	joe := someWork("joe ")
	ana := someWork("ana ")
	jeff := someWork("jeff ")

	t := time.After(5 * time.Second)
	for {
		select {
		case j := <-joe:
			fmt.Println(j)
		case a := <-ana:
			fmt.Println(a)
		case je := <-jeff:
			fmt.Println(je)
		case <-t:
			fmt.Println("I am quitting, its already 5 sec")
			return
		}
	}
}
