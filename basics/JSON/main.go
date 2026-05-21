package main

import (
	"encoding/json"
	"fmt"
)

type course struct {

	// Its optional to use JSON `` and mostly used to change name and not to display password
	Name  string
	Price int
	//Using JSON ``
	Platform string `json:"domain"`
	Password string `json:"-"`
	Tags     []string
}

func main() {
	fmt.Println("Welcome to My JSON code")
	Encoded()

}

func Encoded() {
	lco := []course{
		{"React", 199, "react.in", "ak47", []string{"ReactJS", "Next"}},
		{"Angular", 599, "Angl.in", "abc47", []string{"AngularJS", "Docker"}},
		{"Go", 1999, "GO.in", "556k47", nil},
	}

	FinalJson, err := json.MarshalIndent(lco, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(FinalJson))

}
