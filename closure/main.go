package main

import "fmt"

func intSeq() func() int {
	fmt.Println("intSeq initialize")
	i := 0
	var addFunc = func() int {
		i++
		return i
	}
	return addFunc
}
func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
}
