package main

import (
	"fmt"
)

func main() {
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	// can assign values without the var keyword inside a function only
	age := 14

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)
	fmt.Println("Age: ", age)
}
