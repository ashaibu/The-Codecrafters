# GO BOOLEAN DATA TYPE
# Boolean Data Type

Boolean = true or false only
# How to declare

Go
var b1 bool = true
var b2 = true
var b3 bool   // default
b4 := true
# Key points

Only 2 values → true or false
Default value = false
Used for conditions (if, loops, checks)
# Example output

Go
fmt.Println(b3) // false



# GO INTERGER DATA TYPES
Integer Types: store whole numbers (no decimals).

1. Signed Integers: can be positive or negative.

2. Unsigned Integers: only non-negative values.

# TIPS
 Choose the type based on the range of values you need. Using a smaller type than needed causes errors (e.g., int8 x = 1000 → overflow).


 # GO FLOAT DATA TYPES

 Float Types: store numbers with decimals (positive or negative).

1. float32 – 32 bits, range: -3.4×10³⁸ to 3.4×10³⁸
2. float64 – 64 bits, range: -1.7×10³⁰⁸ to 1.7×10³⁰⁸ (default type)

Tip: Choose the type based on the size of the number you need to store. Using a smaller type than needed causes overflow (e.g., float32 x = 3.4e+39 → error).



# GO STRINGS DATA TYPES

String Type: stores text (sequence of characters).

Strings must be inside double quotes " ".

Example:

var txt1 string = "Hello!" // typed declaration
var txt2 string            // empty string
txt3 := "World 1"          // shorthand declaration

Output:

Type: string, value: Hello!
Type: string, value:
Type: string, value: World 1

# Tip: 
Use var for explicit typing, or := for shorthand declaration with type inferred.