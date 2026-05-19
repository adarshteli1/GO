package main

import "fmt"

type Person struct {
	name string
	age  int
	f    func(string) string
}

func (p Person) getName() string {
	return p.name
}

func main() {
	var p1 Person = Person{age: 23, name: "Addy"}
	value := p1.getName()
	fmt.Println(value)

}
