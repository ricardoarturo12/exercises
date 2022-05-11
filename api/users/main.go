package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

type Album struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	var users []User

	userJson := `[
	{
        "name": "Test",
        "age": 10,
        "phone": "+2334488"
    },
    {
        "name": "Test1",
        "age": 20,
        "phone": "+595059090"
    }]`

	json.Unmarshal([]byte(userJson), &users)
	fmt.Printf("users: %+v", users[0])

	album := &Album{Id: 1, Name: "test"}

	// struct de golang a bytes / para luego pasar a string
	valueByte, _ := json.Marshal(album)

	fmt.Println(string(valueByte))

	// genera un nuevo album
	var s2 string = `
	{"id": 3, "name": "Ricardo"}
	`
	var stuAlb = &Album{}

	err := json.Unmarshal([]byte(s2), stuAlb)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v \n", *stuAlb)
}
