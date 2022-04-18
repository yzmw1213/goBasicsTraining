package main

import (
	"fmt"
)

type point struct {
	x, y int
}

var p = func(passage string, format string, p ...interface{}) {
	fmt.Println(passage)
	fmt.Printf(format, p)
}

func main() {
	p := point{1, 2}

	// %v デフォルトのフォーマット
	fmt.Printf("%v\n", p)

	// %+v 構造体で用いると、フィールド名も加わる
	fmt.Printf("%+v\n", p)

	// %#v Go-syntax representation of the value
	// 値を生成するソース
	fmt.Printf("%#v\n", p)

	// %T 値の型
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", map[string]string{"foo": "Foo", "bar": "Bar"})
	fmt.Printf("%T\n", true)

	//  %t 値がtrueであるかfalseであるか
	fmt.Printf("%t\n", true)

	// %b 数字を2進数に変換
	fmt.Printf("%b\n", 10)
	// %x 数字を16進数に変換(a-fはlower-case)
	fmt.Printf("%x\n", 1010)

	const name, id = "bueller", 17
	// Errof フォーマットを行い、 errorを満たす文字列を返す
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())
}
