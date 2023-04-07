package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	title           string
	original_author string
	num_pages       int
	Language        string `json:"printed_language"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the Server!")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	b1 := Book{
		title:           "C language",
		original_author: "Brian Kernighan and Dennis Ritchie",
		num_pages:       249,
		Language:        "English",
	}

	data, err := json.Marshal(b1)

	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(data))
}

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	b1 := Book{
		title:           "C language",
		original_author: "Brian Kernighan and Dennis Ritchie",
		num_pages:       249,
		Language:        "English",
	}

	b2 := Book{
		title:           "Java",
		original_author: "James Gosling",
		num_pages:       500,
		Language:        "English",
	}

	b3 := Book{
		title:           "C++",
		original_author: "Bjarne Stroustrup",
		num_pages:       350,
		Language:        "English",
	}

	books := []Book{b1, b2, b3}

	data, err := json.MarshalIndent(books, "", "   ")

	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(data))
}

func main() {

	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/get-book", GetBook)
	http.HandleFunc("/get-all-book", GetAllBook)
	log.Println("Starting the server on port 9494")
	log.Fatal(http.ListenAndServe(":9494", nil))
}
