package main

import (
	"fmt"
	"strconv"
)

var p = fmt.Println

// 文字列を数値にパースする
func main() {
	f, _ := strconv.ParseFloat("2.23", 64)
	p(f)

	i, _ := strconv.ParseInt("123", 10, 64)
	p(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	p("ParseIntに16進文字列を与える", d)

	u, _ := strconv.ParseUint("2", 0, 64)
	p("ParseUint is like ParseInt but for unsigned number", u)

	k, _ := strconv.Atoi("135")
	p("Atoi is like ParseInt with base is 10", k)

	s := strconv.Itoa(111222)
	p("Itoa は10進数を文字列に変換する", s)

}
