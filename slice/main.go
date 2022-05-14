package main

import (
	"fmt"
)

func main() {
	length := 3
	capacity := 10
	s := make([]string, length, capacity)
	fmt.Printf("s initialize %v\n", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Printf("s after set %v\n", s)

	s = append(s, "d", "e", "f")
	fmt.Printf("s after append %v\n", s)

	// sliceのコピー
	c := make([]string, len(s))
	copy(c, s)
	fmt.Printf("c after copy %v\n", c)

	// sliceの要素を、範囲指定して取り出す
	// s[start:end] => 実際に取り出される要素の終了位置は、end-1となる
	// インデックス2,3,4の範囲 を切り取る
	l := s[2:5]
	fmt.Printf("l %v\n", l)

	// インデックス0 ~ 4の範囲 を切り取る
	l = s[:5]
	fmt.Printf("l %v\n", l)

	// インデックス1 ~ len(slice) -1 の範囲 を切り取る
	l = s[2:]
	fmt.Printf("l %v\n", l)
	m := s[1:]
	fmt.Printf("m %v\n", m)

}
