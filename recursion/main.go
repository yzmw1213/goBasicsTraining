package main

import "fmt"

// 再起関数
func fact(n int) int {
	// 無限呼び出しを回避するため、再起関数ではベースケースに対してreturnする処理を入れる。
	// n == 0のケースをコメントアウトしプログラムを実行すると、スタックオーバーフローが発生する
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
