package main

import (
	"fmt"
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	i := 20
	for i > 0 {
		// 0-100の数字をランダムに出力する
		fmt.Println(getRandomNumber(100))
		i--
	}
	fmt.Println(getRandomString(30))
}

// getRandomNumber
func getRandomNumber(maxNum int) int {
	// 生成される乱数を毎回ランダムにするため、呼び出すたびに変化する値を乱数シードに設定する
	rand.Seed(time.Now().UnixNano())
	if maxNum <= 0 {
		return maxNum
	}
	return rand.Intn(maxNum)
}

// getRandomString 指定された文字数のランダム文字列を生成
func getRandomString(n int) string {
	// 0以下が指定された場合は空文字を返し、処理を終了する
	if n <= 0 {
		return ""
	}
	// buf 長さnのruneスライス
	buf := make([]rune, n)
	// 元となるruneのスライス
	all := []rune(letters)
	fmt.Println("all", all)
	// bufの各要素に、allの要素からランダムに取得した値を設定する
	for i := range buf {
		buf[i] = all[getRandomNumber(len(all))]
	}
	return string(buf)
}
