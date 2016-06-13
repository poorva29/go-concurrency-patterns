// +build OMIT

// START1 OMIT
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// STOP1 OMIT

// START2 OMIT
func main() {
	go iterator("iterator!") // HL
}

// STOP2 OMIT

func iterator(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
