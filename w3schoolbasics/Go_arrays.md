# GO ARRAYS
Arrays: store multiple values of the same type in a single variable. Length is fixed.

1. Declare an array:

With var:
var arr = [3]int{1,2,3}        // length defined
var arr = [...]int{1,2,3}      // length inferred
With :=:
arr := [5]int{4,5,6,7,8}       // length defined
arr := [...]int{4,5,6,7,8}     // length inferred

2. Access elements:

prices := [3]int{10,20,30}
fmt.Println(prices[0]) // 10
fmt.Println(prices[2]) // 30

3. Change elements:

prices[2] = 50
fmt.Println(prices) // [10 20 50]

4. Default values:

int → 0, string → ""
arr := [5]int{} // [0 0 0 0 0]

5. Initialize specific elements:

arr := [5]int{1:10, 2:40} // [0 10 40 0 0]

6. Find length:

fmt.Println(len(arr)) 

# Tip:
 Array indexes start at 0.