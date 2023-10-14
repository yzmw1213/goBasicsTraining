package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	readFile()
}

func readFile() (err error) {
	file, err := os.Open("text.txt")
	defer func() {
		e := file.Close()
		if e != nil {
			fmt.Println(e)
		}
	}()
	if err != nil {
		return
	}
	data, err := io.ReadAll(file)
	fmt.Printf(string(data))
	return
}
