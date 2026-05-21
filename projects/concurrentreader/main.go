package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	Read("file1.txt")
	// files := []string{
	// 	"file1.txt",
	// 	"file2.txt",
	// 	"file3.txt",
	// }

	// for _, file := range files {
	// 	file, err := os.Open(file)
	// 	if err != nil {
	// 		fmt.Println("Error ", err)
	// 	} else {
	// 		fmt.Println("All Good")
	// 	}

	// 	Read(file)
	// }
}

func Read(file string) {
	charCount := 0
	wordCount := 0
	count := 0
	openedFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		scanner := bufio.NewScanner(openedFile)
		for scanner.Scan() {
			line := scanner.Text()
			count++
			wordCount += len(strings.Fields(line))
			charCount += len([]rune(line))
		}
		fmt.Println("Lines are :", count)
		fmt.Println("Characters are :", charCount)
		fmt.Println("Words are :", wordCount)
	}

}

type Reader struct {
	Lines int
	Words int
	Char  int
}
