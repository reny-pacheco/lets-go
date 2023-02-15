package main

import (
	"fmt"
)

// Source: https://www.digitalocean.com/community/tutorials/how-to-construct-for-loops-in-go

func main() {
	// with initialization
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// without initialization
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}
	// using a break statement
	for num := 0; num < 5; num++ {
		if num == 2 {
			fmt.Println("Number is 2, breaking...")
			break
		}
		fmt.Println(num)
	}

	// iterating an slice using the ForClause
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

	for i := 0; i < len(sharks); i++ {
		fmt.Println(sharks[i])
	}

	// iteratinga an slice using the RangeClause
	for i, shark := range sharks {
		fmt.Println(i, shark)
	}

	// using a blank identifier (_)
	for _, shark := range sharks {
		fmt.Println(shark)
	}

	// adding items to list using range
	for range sharks {
		sharks = append(sharks, "shark")
	}
	fmt.Printf("%q\n", sharks)

	// fill values to a slice using range
	integers := make([]int, 10)
	fmt.Println(integers)

	for i := range integers {
		integers[i] = i
	}
	fmt.Println(integers)

	// iterate through each character in a string
	sammy := "Reny"

	for _, letter := range sammy {
		fmt.Printf("%c\n", letter)
	}

	// iterating a map will return both key and value
	// Note: The order in which a map returns is random
	sammyShark := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}

	for key, value := range sammyShark {
		fmt.Println(key + ": " + value)
	}
}
