package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// type Users struct {
// 	Users []User `json:"users"`
// }

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

func main() {

	// slice de usuarios
	var users []User

	//// abre el archivo y carga en jsonFile
	// jsonFile, err := os.Open("users.json")

	// defer jsonFile.Close()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// byteValue, _ := ioutil.ReadAll(jsonFile)

	// el json file lee como byte para cargar en la estructura
	byteValue, err := ioutil.ReadFile("user.json")

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	json.Unmarshal(byteValue, &users)

	fmt.Printf("%+v", users)

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i].Name)
		fmt.Println(users[i].Age)
		fmt.Println(users[i].Phone)

	}

	file, _ := json.MarshalIndent(users, " ", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}
