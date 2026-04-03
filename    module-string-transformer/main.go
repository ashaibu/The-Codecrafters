package main

import (
	"fmt"
	"strings"
)
func capEachWord(word string) string {
	return strings.ToUpper(word) 
}

func lower(word string) string {
	return strings.ToLower(word)
}

func cap(word string) string {
	return strings.Title(strings.ToLower(word))
}

func Snake(text string) string {
	text = strings.ToLower(text)
	result := ""

	for i := 0; i < len(text); i++ {
		char := text[i]

		if char >= 'a' && char <= 'z' || char >= '0' && char <= '9' {
			result += string(char)
		} else if char == ' ' {
			result += "_"
		}
	}

	return result
}

func Reverse(word string) string {
	s := strings.Fields(word)
	
	for i, j := range s {
		x := []rune (j)
		for a,b := 0, len(x)-1; a < b; a,b = a+1, b-1 {
		x[a], x[b] = x[b], x[a]
	}
	s[i] = string(x)
}
    output := strings.Join(s, " ")
	return output
}
func main() {

	fmt.Println(capEachWord("sentinel is online"))
	fmt.Println(lower("ALERT LEVEL FIVE DETECTED"))
	fmt.Println(cap("director adaeze okonkwo"))
	fmt.Println(cap("THREAT LEVEL elevated"))
	fmt.Println(Reverse("Lagos Nigeria"))
	fmt.Println(Snake("Operation Gopher Protocol"))
	fmt.Println(Snake("Alert! Level 5 detected"))
}