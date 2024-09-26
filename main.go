package main

import (
	"fmt"
	"gin-grpc-generator/parser"
	"os"
)

func main() {
	fmt.Println("RPC Generator for Go")
	filename := "test.proto"
	if !checkFileExist(filename) {
		fmt.Println("File not found")
		return
	}
	data := LoadFile(filename)
	parser.GetService(data)
}

func checkFileExist(filePath string) bool {
	if _, err := os.Stat("source/" + filePath); err != nil {
		return false
	}
	return true
}

func LoadFile(filePath string) string {
	file, err := os.Open("source/" + filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return ""
	}

	size := fi.Size()
	data := make([]byte, size)
	_, err = file.Read(data)
	if err != nil {
		return ""
	}

	return string(data)
}
