package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Books struct {
	ListBook []Book `json:"books"`
}

type Book struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Publisher Publisher `json:"publisher"`
}

type Publisher struct {
	By      string `json:"by"`
	Edition int    `json:"edition"`
}

func main() {

	data := &Books{}

	booksJson, err := ioutil.ReadFile("books.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(booksJson, data)

	if err != nil {
		panic(err)
	}

	// fmt.Printf("%+v", data.ListBook)

	for i := 0; i < len(data.ListBook); i++ {
		fmt.Printf("id %+v ", data.ListBook[i].Id)
		fmt.Printf("title %+v ", data.ListBook[i].Title)
		fmt.Printf("Publisher %+v \n", data.ListBook[i].Publisher.By)
	}

	file, _ := json.MarshalIndent(data, "", "	")
	_ = ioutil.WriteFile("result.json", file, 0644)
}
