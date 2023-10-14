package main

import "bytes"

func main() {
	var b bytes.Buffer
	b.Write([]byte("Hello"))
	plain := make([]byte, 10)
	b.Read(plain)
}
