package main

import (
	"fmt"
	// "strings"
)

func main() {
	for {
	start:
		var num1 float64
		var num2 float64
		var op string

		fmt.Println("1st num")
		_, err1 := fmt.Scan(&num1)

		fmt.Println("2nd num")
		_, err2 := fmt.Scan(&num2)

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid Input")
			goto start
		}

		fmt.Println("choose an operation:\n + for addition\n * for multiplication\n - for subtraction\n  / for division\n or help or quit to exit")
		fmt.Scanln(&op)

		if op == "+" {
			fmt.Println("The result is: ")
			fmt.Println(num1 + num2)
		}

		if op == "*" {
			fmt.Println("The result is: ", num1*num2)
		}
		if op == "-" {
			fmt.Println("This result is: ", num1-num2)

		}
		if op == "/" {
			if num2 == 0 {
				fmt.Println("Can't divide by 0")
				goto start
			}
			fmt.Println("The result is", num1/num2)
		}

		if op == "quit" {
			fmt.Println("Goodbye and thanks for your time")
			break
		}

		if op == "help" {
			fmt.Println("select + for addition\n * for multiplication\n - for subtraction\n  / for division\n or help or quit to exit")
			goto start
		}

	}
}

// omommo
