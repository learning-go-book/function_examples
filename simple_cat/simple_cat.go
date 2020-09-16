package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open(os.Args[1])
	defer func() {
		f.Close()
	}()
	data := make([]byte, 2048)
	var err error
	for err == nil {
		var count int
		count, err = f.Read(data)
		fmt.Print(string(data[:count]))
	}
}
