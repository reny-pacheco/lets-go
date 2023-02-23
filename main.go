package main

import (
	"fmt"
)

func myMessage() {
	fmt.Println("I just got executed!")
}

func funcStr(name string) string {
	return "Hello Guys"
}

func addNums(num1, num2 int) int {
	return num1 + num2
}

func showName(name string) string {
	return "Hello"
}

// declaring a type
type IZ int

// declaring multiple types
type (
	IT  int
	STR string
)

var name STR = "Reny"
var amount IT = 1200

// using constants as enums
const (
	Male   = 0
	Female = 1
)

// auto increment in enums
const (
	a = iota
	b
	c
)

func main() {
	myMessage() // call the function
	fmt.Println(showName("hello"))
	fmt.Println(funcStr("hello"))
	fmt.Println(addNums(2, 3))
	fmt.Println(name, amount)
	fmt.Println(Male)
	fmt.Println(b)
}
