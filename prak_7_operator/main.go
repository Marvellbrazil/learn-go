package main

import "fmt"

func main() {
	x := 5
	y := 2
	
	// Arithmetic operators
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)
	
	// Assignment operators
	var z = "Cogito ergo sum"
	x -= 2
	y += 3
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	
	// Comparison operators
	fmt.Println(x > y)
	fmt.Println(x >= y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x == y)
	fmt.Println(x != y)
	
	// Logical operators
	if x == 3 && y == 5 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	
	if x == 3 || y == 5 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	
	if !(x == 3 && y == 5) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}