package main

import (
	"fmt"
	"time"
)

func selectTimeout() {
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
