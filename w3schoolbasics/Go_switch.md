# GO SWITCH STATEMENT

# Switch Statement
Used to select one of many code blocks to execute.
Go’s switch does not need break—only the matched case runs.
Optional default case runs if no case matches.

# Single-Case switch Syntax

switch expression {
case x:
    // code block if expression == x
case y:
    // code block if expression == y
case z:
    // code block if expression == z
default:
    // code block if no case matches
}

How it works:

Expression is evaluated once.
Its value is compared with each case.
Only the first matched case executes.
If no match, default executes (if present).
Example
day := 4

switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")
case 4:
    fmt.Println("Thursday")
case 5:
    fmt.Println("Friday")
case 6:
    fmt.Println("Saturday")
case 7:
    fmt.Println("Sunday")
}

Output:

Thursday

# Tip:

switch is cleaner than multiple if else if chains when checking one value against many options.

# Go Multi-case switch Statement
# Multiples -Case switch Syntax

Multi-Case switch
You can list multiple values in a single case.
The code block runs if the expression matches any of the values.
Works with an optional default for unmatched values.

Syntax:

switch expression {
case x, y:
    // code runs if expression == x or y
case v, w:
    // code runs if expression == v or w
case z:
    // code runs if expression == z
default:
    // code runs if expression matches none of the cases
}
Example
day := 5

switch day {
case 1, 3, 5:
    fmt.Println("Odd weekday")
case 2, 4:
    fmt.Println("Even weekday")
case 6, 7:
    fmt.Println("Weekend")
default:
    fmt.Println("Invalid day number")
}

Output:

Odd weekday

# Tip:

Multi-case switch is useful when different values need the same action, avoiding repetitive code.

