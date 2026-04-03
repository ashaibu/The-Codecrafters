# GO OUTPUT
Go has 3 ways to print output:
# 1. Print()
Prints without new line
Go
fmt.Print("Hello")
fmt.Print("World")
# Output:
HelloWorld
 2. Println()
Prints with new line
Go
fmt.Println("Hello")
fmt.Println("World")
# Output:

Hello
World
# 3. Printf()
Prints formatted output
Go
fmt.Printf("My age is %d", 20)

#Tip
\n = new line in Print()
# In one line
->  Print = same line
-> Println = new line
-> Printf = formatted output 


# FORMATTING VERBS

# Printf uses verbs (%) to control how output looks.
# General verbs
%v → default value
%T → type
%% → prints %
Go
fmt.Printf("%v %T", 10, 10) // 10 int
🔢 Integers
%d → normal number
%b → binary
%x / %X → hex
%04d → padded (e.g. 0015)
Go
fmt.Printf("%b", 10) // 1010
# 🔤 Strings
%s → plain text
%q → quoted text
%8s → padded text
Go
fmt.Printf("%q", "Hi") // "Hi"
# Boolean
%t → true / false
Go
fmt.Printf("%t", true) // true
# Floats
%f → decimal
%.2f → 2 decimal places
%e → scientific
Go
fmt.Printf("%.2f", 3.141) // 3.14
