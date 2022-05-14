package main

import "fmt"

// intSeqは、エンクロージャ
func intSeq() func() int {
	fmt.Println("intSeq initialize")
	i := 0
	// addFuncは、外側の変数iを参照する。クロージャである。
	// クロージャを生成した事で、iがaddFunc関数内に記憶される。
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
