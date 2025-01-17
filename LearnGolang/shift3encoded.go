package main

import (
	"fmt"
)

func shift3Encode(str string) string {
	shift := 3
	encodedStr := ""
	for _, char := range str {
		shiftedChar := rune(int(char) + shift)
		encodedStr += string(shiftedChar)
	}
	return encodedStr
}

func main() {
	text := "golang encode first program"
	encodedText := shift3Encode(text)
	fmt.Println(encodedText)
}
