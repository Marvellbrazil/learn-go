package main

import (
	"fmt"
)

func main() {
	current_temp := 32
	sweet_temp_in_celcius := 24
	if current_temp > sweet_temp_in_celcius {
		fmt.Println("The temperature is hot")
	}

	time := 20
	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	a := 14
	b := 14
	if a < b {
		fmt.Println("a is less than b.")
	} else if a > b {
		fmt.Println("a is more than b.")
	} else {
		fmt.Println("a and b are equal.")
	}
}

// source from W3Schools
