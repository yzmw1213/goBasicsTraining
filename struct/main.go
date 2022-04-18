package main

import "fmt"

type Person struct {
	name string
	profile map[string]interface{}
}
func main() {
	var Trump Person
	Trump.name= "donald"
	var profile = map[string]interface{}{}

	profile["age"] = 20
	profile["height"] = 150
	profile["hobby"] ="guitar"
	Trump.profile = profile
	fmt.Println(Trump)
}
