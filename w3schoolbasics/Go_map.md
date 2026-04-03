# GO MAP

Map
Stores data in key:value pairs
Unordered, changeable, no duplicate keys
Like a dictionary

1. Create a Map

Using {}

a := map[string]string{
    "brand": "Ford",
    "model": "Mustang",
}

Using make() (recommended)

a := make(map[string]string)
a["brand"] = "Ford"
a["model"] = "Mustang"

2. Access Values
fmt.Println(a["brand"]) // Ford

3. Add or Update
a["year"] = "1964"  // add
a["year"] = "1970"  // update

4. Delete Element
delete(a, "year")

5. Check if Key Exists
val, ok := a["brand"]
fmt.Println(val, ok) // Ford true

6. Loop Through Map
for k, v := range a {
    fmt.Println(k, v)
}
Key Points
Keys must support == (e.g. string, int)
Values can be any type
Maps are unordered
Use make() to safely create empty maps
Maps are reference types (changes affect all copies)