package main

import "fmt"

func main() {
	fmt.Println("The list of Numbers from 1 to 100")
	for i := 1; i <= 100; i++ {
		fmt.Print(i, " ")
	}
}
