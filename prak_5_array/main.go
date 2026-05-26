package main

import "fmt"

func main() {
	// Array must be declared by this following syntax
	// var varname = [length]type{value1, value2, ...}
	var array1 = [3]int{1, 2, 3}
	array2 := [3]string{"Home", "Work", "School"}

	fmt.Println(array1)
	fmt.Println(array2)

	// Or if the length is dynamic, we can infer it by change the length value into ...
	var array3 = [...]int{1, 2, 3}
	array4 := [...]string{"HTML", "CSS", "JS", "TS"}

	fmt.Println(array3)
	fmt.Println(array4)
	
	// To access an element of an array we can use varname[index] like usual
	// And remember, array index is always starting from 0
	// Let's use array4 for example
	fmt.Println(array4[0]) // The output will be 'HTML'
	fmt.Println(array4[1]) // The output will be 'CSS'
	fmt.Println(array4[3]) // The output will be 'TS'
	
	// FYI: We can change an element of an array too by using this syntax
	// For example, we use array 2 index 1
	fmt.Println(array2[1])
	array2[1] = "Cafe"
	fmt.Println(array2[1])
	
	// Let's move on to the array initialization
	array5 := [5]int{} // not init'ed
	array6 := [5]int{1, 2, 3} // partially init'ed
	array7 := [5]int{1, 2, 3, 4, 5} // fully init'ed
	
	// For the uninit'ed elements will be replaced by the default value of the array type
	fmt.Println(array5)
	fmt.Println(array6)
	fmt.Println(array7)
	
	// We can init'ed specific elements too, like this one
	array8 := [5]int{1:6, 3:7}
	fmt.Println(array8)
	
	// To know a length of an array, we can use len()
	// array3 for example, since it's length is inferred
	fmt.Println(len(array3))
}
