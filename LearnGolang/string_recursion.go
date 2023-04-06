// GOLANG PROGRAM TO REVERSE A STRING USING RECURSION
package main

import "fmt"

func reversestring(input string) {
	if len(input) == 0 {
		return
	}

	reversestring(input[1:])
	fmt.Print(string(input[0]))
}

func main() {
	fmt.Println("GOLANG PROGRAM")
	fmt.Println("Entered sentence =")
	var name string
	name = "Caleb Doxsey"
	fmt.Println(name)
	reversestring(name)
}
