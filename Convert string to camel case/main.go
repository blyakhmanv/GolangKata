package main

import "fmt"
import "strings"

func main() {

	fmt.Println(ToCamelCase("the-stealth-warrior"))
	fmt.Println(ToCamelCase("The_Stealth_Warrior"))
}

//ToCamelCase - Kata function
func ToCamelCase(s string) string {
	var ss []string

	if strings.Contains(s, "-") {
		ss = strings.Split(s, "-")
	} else {
		ss = strings.Split(s, "_")
	}

	for i := 1; i < len(ss); i++ {
		ss[i] = strings.Title(ss[i])
	}

	return strings.Join(ss, "")
}
