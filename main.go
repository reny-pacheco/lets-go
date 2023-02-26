package main

import "fmt"

// intro to pointers
func pointers() {
	var creature string = "shark"
	var pointer *string = &creature

	fmt.Println("creature =", creature)
	fmt.Println("pointer =", pointer)

	// dereference the pointer
	fmt.Println("*poiner", *pointer)

	// modify the value stored at the pointer
	*pointer = "jellyfish"
	fmt.Println("*poiner", *pointer)

	// value of creature is also modified
	fmt.Println(creature)
}

// function pointer receivers -------------------
type Creature struct {
	Species string
}

// --> Pass by value
func changeCreature(creature Creature) {
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}

// -----------------------------

func main() {
	// pointers()
	var creature Creature = Creature{Species: "shark"}

	fmt.Printf("1) %+v\n", creature)
	changeCreature(creature)
	fmt.Printf("3) %+v\n", creature)
}
