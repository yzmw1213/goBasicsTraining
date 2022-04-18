package main

import "fmt"

func main () {
	m := make(map[string]string)
	m["company_name"] = "sato"
	m["address"] = "Tokyo"
	fmt.Println(m)

	fmt.Printf("company name is %s\n", m["company_name"])
	fmt.Printf("address is %s\n", m["address"])
	delete(m, "address")

	fmt.Printf("has address after delete key is %t\n", m["address"] !="" )
}