package main

import (
	"fmt"
)

// else statement should not be in a next line as it throws an error
func main() {
	time := 20
	if time < 18 {
		fmt.Println("Good day")
	} else {
		fmt.Println("Good evening")
	}
}
