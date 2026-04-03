# GO CONDITIONS

1. Conditions
A condition evaluates to true or false.
Use comparison operators: <, <=, >, >=, ==, !=
Use logical operators: &&, ||, !

Examples

x > y
x != y
(x > y) && (y > z)
(x == y) || z
2. if Statement

Executes a block if a condition is true:

if x > y {
    fmt.Println("x is greater than y")
}
3. else Statement

Executes a block if the if condition is false:

if x > y {
    fmt.Println("x > y")
} else {
    fmt.Println("x <= y")
}
4. else if Statement

Checks a new condition if the first if is false:

if x > y {
    fmt.Println("x > y")
} else if x == y {
    fmt.Println("x equals y")
} else {
    fmt.Println("x < y")
}
5. switch Statement

Used for multiple alternatives:

switch day {
case "Monday":
    fmt.Println("Start of the week")
case "Friday":
    fmt.Println("Almost weekend")
default:
    fmt.Println("Midweek day")
}

# THE "IF" STATEMENT

# if Statement
Executes a block of code only if a condition is true.
Syntax (note lowercase if):
if condition {
    // code to run if condition is true
}
Examples

1. Using literal values

if 20 > 18 {
    fmt.Println("20 is greater than 18")
}
// Output: 20 is greater than 18

2. Using variables

x := 20
y := 18
if x > y {
    fmt.Println("x is greater than y")
}
// Output: x is greater than y

# Tip:

Uppercase If or IF will cause an error in Go.
Conditions must evaluate to true or false.

# GO IF ELSE STATEMENT

# else Statement
Used to run a block of code if the if condition is false.
Syntax (note the bracket placement: } else {):
if condition {
    // code if condition is true
} else {
    // code if condition is false
}
Examples

1. Using time variable

time := 20
if time < 18 {
    fmt.Println("Good day.")
} else {
    fmt.Println("Good evening.")
}
// Output: Good evening.

2. Using temperature variable

temperature := 14
if temperature > 15 {
    fmt.Println("It is warm out there.")
} else {
    fmt.Println("It is cold out there.")
}
// Output: It is cold out there.
Important Tip
The else must be on the same line as the closing } of if:

✅ Correct:

if condition {
    // code
} else {
    // code
}

❌ Wrong (causes error):

if condition {
    // code
} 
else { // ❌ This raises a syntax error
    // code
}


# ELSE IF STATEMENT 

# else if Statement
Used to check a new condition if the first if condition is false.
Can be combined with if and else to handle multiple conditions.

Syntax:

if condition1 {
    // code if condition1 is true
} else if condition2 {
    // code if condition1 is false and condition2 is true
} else {
    // code if both conditions are false
}
Examples

1. Using time variable

time := 22
if time < 10 {
    fmt.Println("Good morning.")
} else if time < 20 {
    fmt.Println("Good day.")
} else {
    fmt.Println("Good evening.")
}
// Output: Good evening.

2. Comparing two numbers

a := 14
b := 14
if a < b {
    fmt.Println("a is less than b.")
} else if a > b {
    fmt.Println("a is more than b.")
} else {
    fmt.Println("a and b are equal.")
}
// Output: a and b are equal.

3. Multiple conditions – only the first true executes

x := 30
if x >= 10 {
    fmt.Println("x is larger than or equal to 10.")
} else if x > 20 {
    fmt.Println("x is larger than 20.")
} else {
    fmt.Println("x is less than 10.")
}
// Output: x is larger than or equal to 10.

# Tip:

Only the first true condition executes.
Conditions are checked top-down.

# THE NESTEDN IF
 # the nested if statement

 You can put an if statement inside another if statement.
The inner if only runs if the outer if is true.

Syntax:

if condition1 {
    // code if condition1 is true
    if condition2 {
        // code if both condition1 and condition2 are true
    }
}
Example
num := 20

if num >= 10 {
    fmt.Println("Num is more than 10.")
    if num > 15 {
        fmt.Println("Num is also more than 15.")
    }
} else {
    fmt.Println("Num is less than 10.")
}

Output:

Num is more than 10.
Num is also more than 15.

# Tip:

Nested if helps check multiple conditions in hierarchy.
You can nest if statements as deep as needed, but too many levels can make code hard to read.