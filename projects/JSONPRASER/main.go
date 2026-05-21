package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	input()
}
func input() {
	jsonData := `{
		"name":"Jhon",
		"age":26,
		"skills":["CSS", "Next"]
	}`
	var data Data
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error ", err)
	} else {
		fmt.Println("Name is : ", data.Name)
		fmt.Println("Age is : ", data.Age)
		fmt.Println("Skills are  : ", data.Skills)
	}
}

type Data struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Skills []string `json:"skills"`
}
