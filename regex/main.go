package main

import (
	"fmt"
	"regexp"
)

var (
	p       = fmt.Println
	pattern = "p([a-z]+)ch"
)

func main() {
	match, _ := regexp.MatchString(pattern, "peach")
	p(match)

	// Compile Regexp構造体を返す
	r, _ := regexp.Compile(pattern)

	p("MatchString", r.MatchString("peach"))

	// FindString 正規表現の s で最初にマッチしたテキストを持つ文字列を返す
	p("FindString", r.FindString("peach a"))

	// FindStringIndex 正規表現の s で最初にマッチしたテキストを持つ文字列の開始、終了インデックスをスライスで返す
	p("FindStringIndex", r.FindStringIndex("peach a"))

	// FindStringSubmatch 全体のパターンにマッチした部分と、サブマッチ(丸括弧で示された部分) でした部分の情報を返す
	p("FindStringSubmatch", r.FindStringSubmatch("peach a"))

	// FindAllString Allがつく場合は、最初(左を始点とする）のマッチだけでなく、入力の全てのマッチ部分を対象とする
	p("FindAllString", r.FindAllString("peach peach pinch", 3))
}
