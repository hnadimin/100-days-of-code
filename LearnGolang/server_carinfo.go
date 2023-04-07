package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Car struct {
	Manufacturers string
	Model         string
	Price         int
	Enginetype    string
	Mbook         string `json:"printed_mbook"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the Server!")
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	b1 := Car{
		Manufacturers: "BMW",
		Model:         "one series",
		Price:         25000,
		Enginetype:    "Manual",
		Mbook:         "English",
	}

	data, err := json.Marshal(b1)

	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(data))
}

func GetAllCar(w http.ResponseWriter, r *http.Request) {
	b1 := Car{
		Manufacturers: "BMW",
		Model:         "two series",
		Price:         28000,
		Enginetype:    "Manual",
		Mbook:         "English",
	}

	b2 := Car{
		Manufacturers: "BMW",
		Model:         "three series",
		Price:         30000,
		Enginetype:    "Manual",
		Mbook:         "English",
	}

	b3 := Car{
		Manufacturers: "BMW",
		Model:         "four series",
		Price:         32000,
		Enginetype:    "automatic",
		Mbook:         "English",
	}

	b4 := Car{
		Manufacturers: "BMW",
		Model:         "five series",
		Price:         40000,
		Enginetype:    "automatic",
		Mbook:         "English",
	}

	cars := []Car{b1, b2, b3, b4}

	data, err := json.MarshalIndent(cars, "", "   ")

	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(data))
}

func main() {

	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/get-car", GetCar)
	http.HandleFunc("/get-all-car", GetAllCar)
	log.Println("Starting the server on port 9494")
	log.Fatal(http.ListenAndServe(":9494", nil))
}
