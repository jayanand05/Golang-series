package main frrfr

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello and welcome to json series of golang")
	EncodeJson()
}

type course struct {
	Name string `json:"coursename"`
	Website string
	Password string `json:"-"`
	Price int
}

func EncodeJson() {
	courses := []course{
		{"Java", "java.com", "abh2767", 888},
		{"React", "react.com", "uh7767", 973},
		{"Python", "python.com", "knh2767", 299},
	}

	databyte, _ := json.MarshalIndent(courses, "", "\t")
	fmt.Printf("%s\n", databyte)
}
