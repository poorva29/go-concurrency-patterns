// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START1 OMIT
func main() {
	c := fanIn(iterator("One"), iterator("Two")) // HL
	for i := 0; i < 10; i++ {
		fmt.Println(<-c) // HL
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
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller. // HL
}

// STOP2 OMIT

// START3 OMIT
func fanIn(input1, input2 <-chan string) <-chan string { // HL
	c := make(chan string)
	go func() {    for { c <- <-input1 } }() // HL
	go func() {	for { c <- <-input2 } }() // HL
	return c
}

// STOP3 OMIT
