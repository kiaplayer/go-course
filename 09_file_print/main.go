package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("File name is not specified!")
	}
	filePath := os.Args[1]

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error while opening file:", err)
	}

	io.Copy(os.Stdout, f)
}
