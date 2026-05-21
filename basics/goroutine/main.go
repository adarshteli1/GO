package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("THe task Number is ", id)

}

func main() {
	for i := 0; i <= 10; i++ {

		go task(i)

	}
	time.Sleep(time.Millisecond * 2)
}
