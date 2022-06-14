package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}
	// Sleepを使わずに、メインゴルーチンとサブゴルーチンを同期する
	wg.Add(1)
	go sayHello()
	wg.Wait()

	fmt.Println("done")
}
