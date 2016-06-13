// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START1 OMIT
func main() {
	one := iterator("One") // HL
	two := iterator("Two") // HL
	for i := 0; i < 5; i++ {
		fmt.Println(<-one)
		fmt.Println(<-two)
	}
	fmt.Println("You both should stop; I'm leaving.")
}

// STOP1 OMIT

// START2 OMIT
func iterator(msg string) <-chan string { // Returns receive-only channel of strings. // HL
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function. // HL
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller. // HL
}

// STOP2 OMIT
