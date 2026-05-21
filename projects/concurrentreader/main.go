package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	for _, file := range files {
		file, err := os.Open(file)
		if err != nil {
			fmt.Println("Error ", err)
		} else {
			fmt.Println("All Good")
		}

		Read(file)
	}
}

func Read(file *os.File) {
	Words := len(file)
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
	}
	fmt.Println("Lines are :", count)
	fmt.Println("Words are :", Words)
}

type Reader struct {
	Lines int
	Words int
	Char  int
}
