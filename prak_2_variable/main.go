package main

import "fmt"

/*
Variable declarated by this syntax
var varname type = value
or
var varname = value

And for the simplicity
to assign variable that the type following the value type is using := (only works inside a function)
example:
varname := value
*/
func main() {
	var name string = "Marvell"
	var age = 17
	gender := "Male"

	// Variable declaration without initial value
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Value assignments
	a = "This is A"
	b = 2
	c = false

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(gender)
	
	// Multiple variable declaration
	var w, x, y, z int = 1, 2, 5, 7
	
	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	
	// Another example, if it is different type
	var p, q = 1, "Khaleed"
	r, s := "Is cool?", true
	
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)
	
	// Declare a constants using const, that it value cannot be reassigned
	const PI float32 = 3.14
	const GOLDENRATIO float32 = 1.618
	
	fmt.Println(PI)
	fmt.Println(GOLDENRATIO)
}
