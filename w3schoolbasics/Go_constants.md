GO CONSTANTS

Constants = values that cannot change
# How to declare
Go
const PI = 3.14
# Key rules
Must assign value immediately
Cannot be changed later ❌
Can be inside or outside functions
Usually written in UPPERCASE
# Types

1. Typed
Go
const A int = 1

2. Untyped (Go decides type)
Go
const A = 1

# Important

Go
const A = 1
A = 2 // ❌ ERROR
# Multiple constants

Go
const (
  A = 1
  B = 3.14
  C = "Hi!"
)
