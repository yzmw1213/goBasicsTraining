package main

import (
	"fmt"
	"time"
)

// Unix epochからの秒数やミリ秒を扱う
func main() {
	now := time.Now()
	//secs := now.Unix()
	//nons := now.UnixNano()
	fmt.Println(now)
}
