package main

import (
	"fmt"
)

type Creature struct {
	Name string
}

func main() {
	c := Creature{
		Name: "Sammy the Shark",
	}
	fmt.Println(c.Name)

	// inline structs
	i := struct {
		Name string
		Type string
	}{
		Name: "Sammy",
		Type: "Shark",
	}
	fmt.Println(i.Name, "the", i.Type)
}
