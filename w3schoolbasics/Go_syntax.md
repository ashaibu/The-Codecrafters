# A Go program has:
package → defines the program (must be main to run)
import → brings in tools like fmt
func main() → where execution starts
statements → instructions like printing
# Example:
Go
package main
import "fmt"

func main() {
    fmt.Println("Hello World!")
}
# Key rules:
Code runs inside main()
Each line is a statement
{ must be on the same line as func
Go adds ; automatically
# Exercise answer
Go
fmt.Println("Hello World!")
# Final code:
Go
package main
import "fmt"

func main() {
    fmt.Println("Hello World!")
}