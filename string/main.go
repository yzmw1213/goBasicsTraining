package main

import (
	"fmt"
	"strings"
	"unicode"
)

var p = fmt.Println

func main() {
	// strings.Compare 2つの文字列を辞書的に比較する
	// パッケージ内部でもCompareは非推奨であり、パフォーマンスを考えると == , <,>を使うべき
	p("Compare", strings.Compare("b", "a"))

	// Index javascriptのindexOfに該当する
	p("Index", strings.Index("b", "abc"))
	p("Index", strings.Index("gdstdsstdsjse", "tds"))
	p("Index", strings.Index("gdshtdsse", "t"))

	// Contains substrがsに含まれるかどうかを判定する
	p("Contains", strings.Contains("しんばし", "ん"))

	// ContainsAny 2つの文字列に重複する文字が少なくとも1つあるかどうか。
	// 空文字の場合はfalseになる
	p("ContainsAny", strings.ContainsAny("japan", "chi"))

	// Count substrで渡された重複しない文字列パターンが何個含まれるか
	// substrが空文字の場合は, sの文字コード数 + 1が返却される
	p("Count", strings.Count("abcabcabc", "abc"))
	p("Count", strings.Count("abcabcabc", ""))

	// EqualFold 2つの文字がユニコードで大文字小文字の区別をしない場合に等しいかどうかを返す
	p("EqualFold", strings.EqualFold("a", "b"))
	p("EqualFold", strings.EqualFold("a", "A"))
	p("EqualFold", strings.EqualFold("あ", "ァ"))

	// Fields 渡された文字列を空文字でsplitした結果をスライスで返す
	p("Fields", strings.Fields("foo bar　　baz"))

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	// FieldsFunc 文字列が関数f(c rune) bool を満たさないポイントで分割し、そのスライスを返す
	p("FieldsFunc", strings.FieldsFunc("foo.bar3#ba...e,z", f))

	// HasPrefix 文字列がprefixで始まるかどうかを判定する
	p("HasPrefix", strings.HasPrefix("Gopher", "C"))
	p("HasPrefix", strings.HasPrefix("Gopher", "Go"))
	p("HasPrefix", strings.HasPrefix("#low cost", "#"))

	// HasSuffix 文字列がsuffixで終わるかどうかを判定する
	p("HasSuffix", strings.HasSuffix("Gopher", "er"))

	// Join 文字列スライスの各要素をsepで結合した結果を文字列で返す
	p("Join", strings.Join([]string{"2021", "08"}, "-"))

	// LastIndex substrが最後に登場するindexを返す。なければ-1
	p("LastIndex", strings.LastIndex("artificial", "i"))

	// LastIndexAny charに含まれるいずれかの文字コードが最後に現れる場所のindexを返す。
	p("LastIndexAny", strings.LastIndexAny("go gopher keep going", "goHome"))

	// LastIndexByte cが最後に現れる場所のindexを返す
	p("LastIndexByte", strings.LastIndexByte("go gopher keep going", 'i'))

	//strings.Map()

	// Repeat 文字列を指定された回数繰り返す
	p("Repeat", strings.Repeat("PSG ", 5))

	// Replace oldをnewに 指定された数だけ置き換える
	//  nがマイナス値の場合、全てが置き換えられる
	p("Replace", strings.Replace("oink oink oink", "", "ky", 2))
	p("Replace", strings.Replace("oink oink oink", "k", "ky", -1))

	// SplitAfter 文字列を, sepの後で区切る。
	p("SplitAfter", strings.SplitAfter("oink,oink,oink", ","))

	// SplitAfterN 文字列を, sepの後で区切る。
	// nは返すスライスの要素数を指定する
	p("SplitAfterN", strings.SplitAfterN("oink,oink,oink", ",", -1))

	//  各単語の頭文字をTitleケースにして返す
	p("Title", strings.Title("ぁぃぅぇぉ　ぁi"))

	// ToTitleSpecial 与えられた文字列を。指定された文字コードのタイトルケースに変換する
	p("ToTitleSpecial", strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))

	// ToLower 大文字を小文字に変換する
	// かな文字の変換は、unicode.SpecialCaseを利用する
	p("ToLower", strings.ToLower("ABCDE"))

	// Trim
}
