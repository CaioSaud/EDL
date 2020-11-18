package main

import (
	"fmt"
	"time"
)

func escreve(n int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(n)
		n = n + 1
	}
}
func main() {
	go escreve(1)
	escreve(10)
}
