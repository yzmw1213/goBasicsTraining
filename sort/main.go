package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int8
}

type People []Person

func main() {
	//var strs []string = []string{"c", "A", "s", "n"}
	//fmt.Println(sort.StringsAreSorted(strs))
	//sort.Strings(strs)
	//fmt.Println("Strings ", strs)

	sortPeople()
}

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	// Age 昇順でソートされる
	//return p[i].Age < p[j].Age
	// Age 降順でソートされる
	return p[i].Age > p[j].Age
}

func (p People) Swap(i, j int) {
	//var iP = p[i]
	//var jP = p[j]
	//p[i] = jP
	//p[j] = iP
	p[i], p[j] = p[j], p[i]
}
func sortPeople() {
	var people People = []Person{
		{"Tom", 16},
		{"Chris", 20},
		{"Jhon", 37},
		{"Bush", 1},
	}
	fmt.Println(people)
	sort.Sort(people)
	fmt.Println(people)
}
