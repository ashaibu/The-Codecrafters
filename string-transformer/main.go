package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// ----- TRANSFORMATION FUNCTIONS -----
func upper(text string) string {
	return strings.ToUpper(text)
}

func lower(text string) string {
	return strings.ToLower(text)
}

func capWords(text string) string {
	words := strings.Fields(text)
	result := ""
	for i, word := range words {
		word = strings.ToLower(word)
		result += strings.ToUpper(word[:1]) + word[1:]
		if i != len(words)-1 {
			result += " "
		}
	}
	return result
}

func title(text string) string {
	smallWords := map[string]bool{
		"a": true, "an": true, "the": true, "and": true, "but": true, "or": true,
		"for": true, "nor": true, "on": true, "at": true, "to": true, "by": true,
		"in": true, "of": true, "up": true, "as": true, "is": true, "it": true,
	}
	words := strings.Fields(text)
	result := ""
	for i, word := range words {
		word = strings.ToLower(word)
		if i == 0 || !smallWords[word] {
			result += strings.ToUpper(word[:1]) + word[1:]
		} else {
			result += word
		}
		if i != len(words)-1 {
			result += " "
		}
	}
	return result
}

func snake(text string) string {
	text = strings.ToLower(text)
	result := ""
	for _, char := range text {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			result += string(char)
		} else if char == ' ' {
			result += "_"
		}
	}
	return result
}

func reverse(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		runes := []rune(w)
		for a, b := 0, len(runes)-1; a < b; a, b = a+1, b-1 {
			runes[a], runes[b] = runes[b], runes[a]
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

// ----- COUNT FUNCTION -----
func countText(text string) string {
	totalChars := len(text)
	totalLetters := 0
	totalSpaces := 0
	for _, c := range text {
		if unicode.IsLetter(c) {
			totalLetters++
		} else if unicode.IsSpace(c) {
			totalSpaces++
		}
	}
	totalWords := len(strings.Fields(text))
	return fmt.Sprintf("Total characters: %d\nTotal letters: %d\nTotal words: %d\nTotal spaces: %d",
		totalChars, totalLetters, totalWords, totalSpaces)
}

// ----- MAIN FUNCTION -----
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	history := []string{}

	for {
		fmt.Print("Enter command: ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		words := strings.Fields(input)
		if len(words) == 0 {
			fmt.Println("✗ No text provided. Usage: command <text>")
			continue
		}

		command := strings.ToLower(words[0])
		text := strings.Join(words[1:], " ")

		if command == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			break
		}

		if command == "history" {
			if len(history) == 0 {
				fmt.Println("No history yet.")
			} else {
				fmt.Println("Last 5 transformations:")
				start := 0
				if len(history) > 5 {
					start = len(history) - 5
				}
				for i := start; i < len(history); i++ {
					fmt.Println(history[i])
				}
			}
			continue
		}

		if text == "" {
			fmt.Println("✗ No text provided. Usage: command <text>")
			continue
		}

		var result string
		switch command {
		case "upper":
			result = upper(text)
		case "lower":
			result = lower(text)
		case "cap":
			result = capWords(text)
		case "title":
			result = title(text)
		case "snake":
			result = snake(text)
		case "reverse":
			result = reverse(text)
		case "count":
			result = countText(text)
		default:
			fmt.Println("✗ Unknown command:", command)
			fmt.Println("Valid commands: upper, lower, cap, title, snake, reverse, count, history, exit")
			continue
		}

		fmt.Println(result)
		history = append(history, fmt.Sprintf("%s: %s -> %s", command, text, result))
	}
}