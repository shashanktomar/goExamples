// A concurrent prime sieve

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func primeSeive() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		print(prime, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	elapsedTime := time.Since(start)
	fmt.Println(elapsedTime)
}

func primeSimple() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	primes := []int{2}
	fmt.Println(2)
	for i := 3; len(primes) < 10; i++ {
		isPrime := true
		for _, p := range primes {
			if i%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
			fmt.Println(i)
		}
	}
	elapsedTime := time.Since(start)
	fmt.Println(elapsedTime)
}

func main() {
	primeSeive()
	primeSimple()
}
