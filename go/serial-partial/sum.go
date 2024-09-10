package main

import (
	"bufio"
	"fmt"
	"os"
)

// read a file from a filepath and return a slice of bytes
func readFile(filePath string) (*os.File, error) {
	data, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

// sum all bytes of a file
func sum(filePath string) ([]int, error) {
	data, err := readFile(filePath)
	var somas []int
	if err != nil {
		return somas, err
	}
	reader := bufio.NewReader(data)
	buffer := make([]byte, 100)

	for {
		n, _ := reader.Read(buffer)
		fmt.Println(string(buffer[:n]))
	}
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}
	for _, path := range os.Args[1:] {
		sum(path)
	}
}
