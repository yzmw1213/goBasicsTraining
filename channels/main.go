package main

import (
	"fmt"
	"time"
)

func main() {
	//channel()
	//bufferdChannel()
	channelSynchronization()
}

func channel() {
	message := make(chan string)

	go func() { message <- "ping"}()
	msg := <- message
	fmt.Println(msg)
}

func bufferdChannel() {
	// 4単語まで受け入れるチャネル
	message := make(chan string, 4)


	message <- "this"
	message <- "is"
 	message <- "bufferd"
	message <- "channel"
	// all goroutines are asleep !
	//message <- "aaa"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)

}


// channelSynchronization チャネルを非同期で実行する
func channelSynchronization() {
	//done 他のゴルチンにworkが完了した事を伝えるチャネル
	done := make(chan bool,1)
	// workerは、doneのreceiver無しでは実行されない。
	go worker(done)
	// receive that worker is done
	// receiverをコメントアウトすると、プログラムはworkerを実行せずに終了する
	//<-done
}

// worker ゴルチン上で動く
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 処理が完了したことを伝える
	done <- true
}