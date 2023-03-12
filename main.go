package main

import (
	"fmt"
	"math"
)

// src: https://www.digitalocean.com/community/tutorials/how-to-use-interfaces-in-go

// type Article struct {
// 	Title  string
// 	Author string
// }

// func (a Article) String() string {
// 	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
// }

// func (a Article) ShowAuthor() {
// 	fmt.Println(a.Author, "Showing author")
// }

// type Stringer interface {
// 	String() string
// 	ShowAuthor()
// }

// func main() {
// 	a := Article{
// 		Title:  "Understanding Interfaces in Go",
// 		Author: "Sammy Shark",
// 	}
// 	Print(a)
// }

// func Print(s Stringer) {
// 	fmt.Println(s.String())
// 	s.ShowAuthor()
// }

// Multiple Behaviors in an Interface

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f}", c.Radius)
}

type Square struct {
	Width  float64
	Height float64
}

func (s Square) Area() float64 {
	return s.Width * s.Height
}

func (s Square) String() string {
	return fmt.Sprintf("Square {Width: %.2f, Height: %.2f}", s.Width, s.Height)
}

type Sizer interface {
	Area() float64
}

type Shaper interface {
	Sizer
	fmt.Stringer
}

func main() {
	c := Circle{Radius: 10}
	PrintArea(c)

	s := Square{Height: 10, Width: 5}
	PrintArea(s)

	l := Less(c, s)
	fmt.Printf("%v is the smallest\n", l)
}

// returns the Sizer interface
func Less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}

func PrintArea(s Shaper) {
	fmt.Printf("area of %s is %.2f\n", s.String(), s.Area())
}
