// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START1 OMIT

func main() {
	iterator("iterator!") // HL
}

// STOP1 OMIT

// START2 OMIT
func iterator(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// STOP2 OMIT
