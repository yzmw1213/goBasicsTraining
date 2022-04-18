package main

import "fmt"

func main() {
	// 送受信する値の型を揃えることで安全性を高める
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan string,msg string) {
	pings <- msg
}

func pong(pings chan string, pongs chan string) {
	// pingsチャネルから値を受信する
	msg := <- pings
	// 受け取った値をpongsチャネルに送信する
	pongs <- msg
}