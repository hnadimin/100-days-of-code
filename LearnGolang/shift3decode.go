package main

import (
	"bufio"
	"fmt"
	"os"
)

func shift3Decode(text string) string {
	var decodedText string
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			decodedChar := rune((int(char-'a')-3+26)%26 + 'a')
			decodedText += string(decodedChar)
		} else {
			decodedText += string(char)
		}
	}
	return decodedText
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to be decoded: ")
	text, _ := reader.ReadString('\n')
	decodedText := shift3Decode(text)
	fmt.Println("Decoded text:", decodedText)
}
