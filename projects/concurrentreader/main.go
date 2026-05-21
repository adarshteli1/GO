package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)

	go Read("file2.txt")
	go Read("file2.txt")
	go Read("file3.txt")

	wg.Wait()

	//The Below COde was Sending File one By One
	/*files := []string{
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
	}*/
}

func Read(file string) {
	defer wg.Done()
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
