package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	fmt.Println(wave("hello"))
}

func wave(words string) []string {
	var result []string
	result = make([]string, 0, len(words))
	for i := 0; i < len(words); i++ {
		if !(unicode.IsSpace(rune(words[i]))) {
			fmt.Println((words[:i+1]))
			result = append(result, words[:i]+strings.ToUpper(words[i:i+1])+words[i+1:])
		}
	}

	return result
}
