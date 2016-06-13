// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	iterator("iterator!")
}

// START OMIT
func iterator(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}

// STOP OMIT
