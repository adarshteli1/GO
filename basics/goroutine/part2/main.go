package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	websites := []string{
		"https://www.google.com/",
		"https://www.github.com/",
		"https://www.instagram.com/",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("OOPS in Endpoint")
	} else {
		fmt.Printf("%d Status code of %s\n", res.StatusCode, endpoint)

	}
}
