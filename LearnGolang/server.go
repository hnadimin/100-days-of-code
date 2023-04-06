package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Title    string
	Author   string
	Pages    int
	Language string `json:"printed_language"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the Server!")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	b1 := Book{
		Title:    "Daffodils",
		Author:   "Wordsworth",
		Pages:    249,
		Language: "English",
	}

	data, err := json.Marshal(b1)

	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(data))
}

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	b1 := Book{
		Title:    "Daffodils",
		Author:   "Wordsworth",
		Pages:    249,
		Language: "English",
	}

	b2 := Book{
		Title:    "Bhagvad Gita",
		Author:   "Ved Vyas",
		Pages:    500,
		Language: "Sanskrit",
	}

	b3 := Book{
		Title:    "Barrister Parvateesam",
		Author:   "Mokkapati Narasimha Sastry",
		Pages:    350,
		Language: "Telugu",
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
