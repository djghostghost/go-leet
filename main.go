package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-t.C:
			fmt.Printf("Hello")
		}

	}
}
