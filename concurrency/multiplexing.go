package main

import "fmt"

func multiplexer() {
	joe := boringGenerator("joe ")
	ana := boringGenerator("ana ")
	club := fanIn(joe, ana)
	for i := 0; i < 10; i++ {
		fmt.Printf("You say: %q\n", <-club)
	}
}

func fanIn(channnels ...<-chan string) <-chan string {
	ch := make(chan string)
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
