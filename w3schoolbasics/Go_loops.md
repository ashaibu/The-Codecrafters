# The for Loop
The only loop in Go.
Used to repeat a block of code a specified number of times.
Each iteration executes the code block once.

Basic Syntax:

for init; condition; post {
    // code to run each iteration
}
init → initialize loop counter
condition → evaluated each loop, loop runs if true
post → executed after each iteration (usually increments counter)
Example 1: Counting 0–4
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

Output:

0
1
2
3
4

# The continue Statement
Skips the current iteration and moves to the next.
for i := 0; i < 5; i++ {
    if i == 3 {
        continue
    }
    fmt.Println(i)
}

Output:

0
1
2
4


# The break Statement
Stops the loop completely.
for i := 0; i < 5; i++ {
    if i == 3 {
        break
    }
    fmt.Println(i)
}

Output:

0
1
2


# Nested Loops
Loop inside another loop.
adj := [2]string{"big", "tasty"}
fruits := [3]string{"apple", "orange", "banana"}

for i := 0; i < len(adj); i++ {
    for j := 0; j < len(fruits); j++ {
        fmt.Println(adj[i], fruits[j])
    }
}

Output:

big apple
big orange
big banana
tasty apple
tasty orange
tasty banana


# The range Keyword
Iterates through arrays, slices, or maps.
Returns index and value.
fruits := [3]string{"apple", "orange", "banana"}
for idx, val := range fruits {
    fmt.Printf("%v\t%v\n", idx, val)
}

Output:

0   apple
1   orange
2   banana

# Omitting index or value:

// Only values
for _, val := range fruits { fmt.Println(val) }

// Only indexes
for idx, _ := range fruits { fmt.Println(idx) }