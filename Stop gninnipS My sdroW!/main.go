package main

import "fmt"
import "strings"

//import "bytes"

func main() {

	fmt.Println(spinWords("Hey fellow warriors"))

}

func spinWords(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words); i++ {
		if len(words[i]) > 5 {
			words[i]= reverse(words[i])
		}
	}

	return strings.Join(words," ")
} // SpinWords

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
