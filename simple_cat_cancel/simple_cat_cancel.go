package main

import (
	"fmt"
	"os"
)

func getFile(name string) (*os.File, func()) {
	file, _ := os.Open(name)
	return file, func() {
		file.Close()
	}
}

func main() {
	f, closer := getFile(os.Args[1])
	defer closer()
	data := make([]byte, 2048)
	var err error
	for err == nil {
		var count int
		count, err = f.Read(data)
		fmt.Print(string(data[:count]))
	}
}
