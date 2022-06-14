package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from+":", i)
	}
}
func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("run")
	// Sleepしないとgoroutineが実行されない??
	// 本来は、Waitgroupを使う
	time.Sleep(time.Second)

	fmt.Println("done")
}
