package main

import "fmt"

func main () {
	// 要素数5, intのスライス
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)

	b := [5]int{1,2,3,4,5}
	fmt.Println(b)
	//b = append(b,19) bは固定長の配列であるため、要素数を変更することはできない。

	var twoD [2][3]int

	for i:= 0; i<2; i++ {
		for j:= 0; j <3;j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println(twoD)

}