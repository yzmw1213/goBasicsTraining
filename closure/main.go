package main

import "fmt"

func intSeq() func() int {
	fmt.Println("intSeq initialize")
	i:= 0
	return func() int {
		i++
		return i
	}
}
func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt :=intSeq()
	fmt.Println(newInt())
}
