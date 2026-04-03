# GO SLICES

Slices: like arrays but dynamic—their length can grow or shrink.

1. Create a slice with values
myslice := []int{1,2,3} // length 3, capacity 3
empty := []int{}         // length 0, capacity 0

fmt.Println(len(myslice)) // 3
fmt.Println(cap(myslice)) // 3
len(slice) → number of elements
cap(slice) → maximum elements slice can hold without reallocating

2. Create a slice from an array
arr := [6]int{10,11,12,13,14,15}
myslice := arr[2:4] // take elements at index 2 and 3
fmt.Println(myslice)       // [12 13]
fmt.Println(len(myslice))  // 2
fmt.Println(cap(myslice))  // 4 (can grow to end of array)

3. Create a slice with make()
s1 := make([]int, 5, 10) // length 5, capacity 10
s2 := make([]int, 5)     // length 5, capacity 5

# Tip:

Slices are more flexible than arrays.
Always check len() and cap() to know size and growth potential.



# MODIFY SLICES
# Go Access, Change, Append and Copy Slices

1. Access elements

Slice indexes start at 0:

prices := []int{10,20,30}
fmt.Println(prices[0]) // 10 (first element)
fmt.Println(prices[2]) // 30 (third element)
2. Change elements

Change a slice element by index:

prices[2] = 50
fmt.Println(prices) // [10 20 50]
3. Append elements

Add new elements to the end with append():

myslice := []int{1,2,3,4,5,6}
myslice = append(myslice, 20, 21)
fmt.Println(myslice) // [1 2 3 4 5 6 20 21]

# Tip:

Appending may increase the slice capacity if needed.
Use len(slice) to check current number of elements, cap(slice) to see how much it can grow without reallocating.