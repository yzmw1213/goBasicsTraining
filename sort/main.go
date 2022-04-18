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
	return p[i].Name < p[j].Name
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
	}
	fmt.Println(people)
	sort.Sort(people)
	fmt.Println(people)
}
