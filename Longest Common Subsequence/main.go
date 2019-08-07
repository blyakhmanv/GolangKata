package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(LCS("132535365", "123456789"))
	fmt.Println(LCS("anothertest", "notatest"))
	fmt.Println(LCS("finaltest", "zzzfinallyzzz"))

	fmt.Println(LCS("dehagfcib", "geacbidhf"))
	fmt.Println(LCS("geacbidhf", "dehagfcib"))
	//fmt.Println(LCS("cadhfiebg", "hfbgedaic"))
	//fmt.Println(LCS("agdfebihc", "cabhgfide"))
}

//LCS - find aubsequence
func LCS(s1, s2 string) string {
	results1 := ""
	results2 := ""
	temp := ""

	for i := 0; i < len(s1); i++ {
		temp = lscsupport(s1[i:], s2)
		if len(temp) > len(results1) {
			results1 = temp
		}
	}

	for i := 0; i < len(s2); i++ {
		temp = lscsupport(s2[i:], s1)
		if len(temp) > len(results2) {
			results2 = temp
		}
	}

	if len(s1) < len(s2) {
		return results1
	}
	if len(s2) < len(s1) {
		return results2
	}

	if len(results1) > len(results2) {
		return results1
	} else {
		return results2
	}
}

func lscsupport(s1, s2 string) string {
	var result strings.Builder
	tempind := 0
	chars := []rune(s1)
	for i := 0; i < len(chars); i++ {
		if strings.Contains(s2, string(chars[i])) {
			if strings.LastIndex(s2, string(chars[i])) >= tempind {
				result.WriteRune(chars[i])
				tempind = strings.Index(s2, string(chars[i]))
			}
		}

	}

	return result.String()

}
