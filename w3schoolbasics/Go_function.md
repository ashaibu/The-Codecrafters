# GO FUNCTIONS
# Create/call function
Function
A block of code that can be reused.
It does not run automatically—it runs only when called.
1. Create (Declare) a Function
Use func keyword
Give it a name
Add code inside {}
func myFunction() {
    // code to run
}
2. Call a Function
Use the function name followed by ()
myFunction()
Example
package main
import ("fmt")

func myMessage() {
    fmt.Println("I just got executed!")
}

func main() {
    myMessage() // calling the function
}

Output:

I just got executed!

# Tip:

Functions help you reuse code and keep programs organized


# Go Function Parameters and Arguments

# Parameters vs Arguments
Parameter → variable in the function definition
Argument → actual value passed to the function

1. Function with One Parameter
func familyName(fname string) {
    fmt.Println("Hello", fname, "Refsnes")
}

Calling the function:

familyName("Liam")
familyName("Jenny")
familyName("Anja")

Output:

Hello Liam Refsnes
Hello Jenny Refsnes
Hello Anja Refsnes

2. Multiple Parameters
You can pass more than one value
func familyName(fname string, age int) {
    fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

Calling:

familyName("Liam", 3)
familyName("Jenny", 14)
familyName("Anja", 30)

Output:

Hello 3 year old Liam Refsnes
Hello 14 year old Jenny Refsnes
Hello 30 year old Anja Refsnes
Key Points
Parameters are defined in the function: fname string
Arguments are the values passed: "Liam"
You can have multiple parameters, separated by commas



# Go Function Returns

# Return Values
Functions can send back (return) a value.
You must specify the return type.

Syntax:

func FunctionName(params) returnType {
    return value
}
1. Basic Return Example
func add(x int, y int) int {
    return x + y
}

fmt.Println(add(1, 2)) // 3


2. Named Return Values
You can name the return variable.
Use return without specifying the variable (called naked return).
func add(x int, y int) (result int) {
    result = x + y
    return
}
3. Store Return Value
total := add(1, 2)
fmt.Println(total) // 3


4. Multiple Return Values
Go functions can return more than one value.
func myFunction(x int, y string) (int, string) {
    return x + x, y + " World!"
}

fmt.Println(myFunction(5, "Hello"))
// Output: 10 Hello World!
Key Points
Use return to send values back
Must define return type(s)
Can return one or multiple values
Can use named return variables




# Go RECURSION FUNCTIONS

Recursion
A function that calls itself.
Must have a stop condition to avoid infinite loops.
How it works
Function calls itself
Each call moves closer to a stop condition
Stops when condition is met

Example 1: Counting
func testcount(x int) int {
    if x == 11 {
        return 0
    }
    fmt.Println(x)
    return testcount(x + 1)
}

Output:

1
2
3
4
5
6
7
8
9
10
Example 2: Factorial
func factorial(x float64) float64 {
    if x > 0 {
        return x * factorial(x-1)
    }
    return 1
}

Call:

fmt.Println(factorial(4)) // 24
Key Points
Must have a base case (stop condition)
Each call should move toward stopping
Useful for problems like factorials, trees, loops
 
 # Warning ⚠️
No stop condition → infinite recursion (crash)
Too many calls → high memory usage
