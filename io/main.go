package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//fmt.Println(writeString("01234567"))
	//Write()
	//ReadAll()
	Copy()
}

func writeString(s string) string {
	file, err := os.Create("data/text.txt")
	if err != nil {
		return ""
	}
	if _, err := io.WriteString(file, s); err != nil {
		return ""
	}
	b, err := ioutil.ReadFile("data/test.txt")
	if err != nil {
		return ""
	}
	return string(b)
}

func Write() (err error) {
	file, _ := os.Create("data/file.txt")
	writer := io.Writer(file)
	n, err := writer.Write([]byte("hello"))
	if err != nil {
		return err
	}
	fmt.Println(n)
	file.Close()
	return
}

func Read() (err error) {
	file, _ := os.Open("data/file.txt")
	reader := io.Reader(file)
	buffer := make([]byte, 10)
	// Readは、bufferの中に、len(buffer)までの読み込みを行う
	n, err := reader.Read(buffer)
	defer func() {
		e := file.Close()
		if e != nil {
			fmt.Errorf("%s", e)
		}
	}()
	if err != nil {
		return err
	}
	fmt.Printf("Read n={%d}, buffer={%v}", n, string(buffer))
	return
}

func ReadAll() (err error) {
	file, _ := os.Open("data/file.txt")
	reader := io.Reader(file)
	buffer, err := io.ReadAll(reader)
	fmt.Printf("Read buffer={%v}", string(buffer))
	seeker := reader.(io.Seeker)
	seeker.Seek(0, io.SeekStart)
	// 2回目の読み込みは行えない
	buffer, err = io.ReadAll(reader)
	fmt.Printf("Read buffer={%v}", string(buffer))
	defer func() {
		e := file.Close()
		if e != nil {
			fmt.Errorf("%s", e)
		}
	}()
	// ReadAll reads from r until an error or EOF
	// ReadAllは、ReaderからEOFまたはエラーが発生するまで全てを読み込む
	if err != nil {
		return err
	}
	return
}

func Copy() (err error) {
	file, _ := os.Open("data/file.txt")
	writer := os.Stdout
	r := strings.NewReader("some io.Reader stream to be read")
	if _, err := io.Copy(writer, r); err != nil {
		fmt.Errorf("%v", err)

		return err
	}
	reader := io.Reader(file)
	buf, err := io.ReadAll(reader)
	defer func() {
		e := file.Close()
		if e != nil {
			fmt.Errorf("%s", e)
		}
	}()
	if err != nil {
		return err
	}
	fmt.Println(string(buf))

	return
}
