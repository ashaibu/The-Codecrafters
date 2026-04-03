# GO STRUCT

Struct (Structure)
A struct groups different data types into one variable.
Used to represent real-world data (records).

1. Declare a Struct
type Person struct {
    name   string
    age    int
    job    string
    salary int
}
2. Create and Use Struct
var p Person

p.name = "Hege"
p.age = 45
p.job = "Teacher"
p.salary = 6000

3. Access Struct Members
Use the dot (.) operator:
fmt.Println(p.name)
fmt.Println(p.age)

4. Example with Two Structs
pers1 := Person{name: "Hege", age: 45, job: "Teacher", salary: 6000}
pers2 := Person{name: "Cecilie", age: 24, job: "Marketing", salary: 4500}

fmt.Println(pers1.name, pers1.age)
fmt.Println(pers2.name, pers2.age)
Key Points
Struct = group of different data types
Use type + struct to define
Access fields with .
Useful for organizing related data