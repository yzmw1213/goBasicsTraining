package main

import (
	"fmt"
	"time"
)

func main() {
	//	default句をもったselectにより、ブロックしない送受信を行える。

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hello"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	go func() {
		time.Sleep(1 * time.Second)
		signals <- true
	}()

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	//	default句をコメントアウトするとsignal受信が行われる
	default:
		fmt.Println("there is no activity")
	}
}
