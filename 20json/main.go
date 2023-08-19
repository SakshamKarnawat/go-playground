package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"` // this field will not be shown in the json encoded data
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("json in golang!")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"Docker", 20, "Udemy", "123", []string{"docker", "container"}},
		{"Puppet", 30, "Udemy", "123", []string{"puppet", "devops"}},
		{"ReactJS", 40, "Udacity", "123", []string{"react", "frontend"}},
		{"NodeJS", 50, "Coursera", "123", []string{"node", "backend"}},
		{"MongoDB", 60, "MongoDB", "123", []string{"mongo", "database"}},
		{"Test", 60, "Test", "123", nil},
	}

	// package this data into json
	finalJson, err := json.MarshalIndent(myCourses, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
	{
		"coursename": "MongoDB",
		"price": 60,
		"platform": "MongoDB",
		"tags": ["mongo", "database"]
	}
	`)

	// create a variable of type course, and then pass the address of this variable to the Unmarshal function to
	// update the value of this variable
	var myCourse course // same: var myCourse = course{}

	checkValid := json.Valid(jsonData)
	if checkValid {
		json.Unmarshal(jsonData, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON is not valid")
	}
}
