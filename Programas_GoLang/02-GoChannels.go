package main

import (
	"fmt"
	"time"
)

func escreve(s int, done chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
		s = s * 2
	}
  //avisando a goroutine principal que finalizou
	done <- "Terminei"
}
func main() {
	var c = 1
  //declaracao gochannel
	done := make(chan string)
	go escreve(1, done)
	for j := 0; j < 5000; j++ {
		fmt.Println(c)
		c = c + 1
	}
  //finalizando a goroutine principal só após o envio do fim da goroutine escreve
	fmt.Println(<-done)
}
