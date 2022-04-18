package main

import "fmt"

// 引数の数が可変である
func sum(nums ...int) {
	fmt.Printf("function sum vars: %v\n", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("total : %v\n", total)
}
func main() {
	sum(1,2,3)
	sum(1,2,3, 4)

	var nums = []int{1,2,3,4,5,6}
	sum(nums...)
}
