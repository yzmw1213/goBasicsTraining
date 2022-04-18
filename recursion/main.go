package main

import "fmt"

// 再起関数
func fact(n int) int {
	if n == 0 { return 1}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
