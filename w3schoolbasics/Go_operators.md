# OPERATORS

1. Addition (+)
Adds two values, a value and a variable, or two variables.
a := 15 + 25
fmt.Println(a) // 40
sum1 := 100 + 50     // 150
sum2 := sum1 + 250   // 400
sum3 := sum2 + sum2  // 800
fmt.Println(sum3)    // 800

# Tip:

+ works for numbers (int, float) and also strings (concatenation).


# ARITHMETICS

| Operator | Name           | Description                      | Example |
| -------- | -------------- | -------------------------------- | ------- |
| `+`      | Addition       | Adds two values                  | `x + y` |
| `-`      | Subtraction    | Subtracts one value from another | `x - y` |
| `*`      | Multiplication | Multiplies two values            | `x * y` |
| `/`      | Division       | Divides one value by another     | `x / y` |
| `%`      | Modulus        | Returns remainder of division    | `x % y` |
| `++`     | Increment      | Increases value by 1             | `x++`   |
| `--`     | Decrement      | Decreases value by 1             | `x--`   |


# Exercise Example

Multiply 10 by 5 and print the result:

package main
import ("fmt")

func main() {
    result := 10 * 5
    fmt.Println(result) // 50
}


# GO ASSIGNMENT OPERATORS

1. Basic Assignment
x := 10  // or var x = 10
fmt.Println(x) // 10
= assigns a value to a variable.

2. Combined Assignment
These operators update the variable by performing an operation first:
Operator	Example	Same As
+=	x += 3	x = x + 3
-=	x -= 3	x = x - 3
*=	x *= 3	x = x * 3
/=	x /= 3	x = x / 3
%=	x %= 3	x = x % 3
&=	x &= 3	x = x & 3
`	=`	`x
^=	x ^= 3	x = x ^ 3
>>=	x >>= 3	x = x >> 3
<<=	x <<= 3	x = x << 3

# Example: Addition Assignment
x := 10
x += 5   // x = x + 5
fmt.Println(x) // 15


# GO COMPARISM OPERATORS

# Comparison Operators

Used to compare two values.
Return type: bool (true or false).
Operator	Name	Example
==	Equal to	x == y
!=	Not equal to	x != y
>	Greater than	x > y
<	Less than	x < y
>=	Greater than or equal to	x >= y
<=	Less than or equal to	x <= y

# Example
x := 5
y := 3
fmt.Println(x > y)  // true
fmt.Println(x == y) // false


# GO LOGICAL OPERATORS

# Logical Operators
Used to combine or invert boolean expressions.
Operator	Name	Description	Example
&&	Logical AND	True if both statements are true	x < 5 && x < 10
`		`	Logical OR
!	Logical NOT	Reverses the result	!(x < 5 && x < 10)

# Example
x := 3
y := 5

fmt.Println(x < 5 && y > 2) // true
fmt.Println(x > 5 || y > 2) // true
fmt.Println(!(x < 5))       // false

# GO BITWISE OPERATORS

# Bitwise Operators
Operate on binary representations of numbers.
Operator	Name	Description	Example
&	AND	1 if both bits are 1	x & y
`	`	OR	1 if at least one bit is 1
^	XOR	1 if only one bit is 1	x ^ y
<<	Left shift	Shift bits left, fill 0 from right	x << 2
>>	Right shift	Shift bits right, fill left with sign bit	x >> 2
Example
x := 5  // 0101 in binary
y := 3  // 0011 in binary

fmt.Println(x & y)  // 1  (0001)
fmt.Println(x | y)  // 7  (0111)
fmt.Println(x ^ y)  // 6  (0110)
fmt.Println(x << 1) // 10 (1010)
fmt.Println(x >> 1) // 2  (0010)