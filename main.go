package main

import (
	"fmt"
	"math"
)

func main() {

	// ex #58
	i, s := bar()
	fmt.Println("#58 foo:", foo(), "bar:", i, s)

	// ex #59
	vi := []int{1, 2, 3, 4, 5}
	y := foovar(vi...)
	z := barvar(vi)
	fmt.Println("#59 foovar:", y, "barvar:", z)

	// ex #60
	// LIFO
	defer fmt.Println("#60 foodefer Print me last")
	defer fmt.Println("#60 bardefer Print me second to last")
	fmt.Println("#60 Print me first...")

	// ex #61
	p := Person{"Karen", 53}
	fmt.Println("#61", p.Speak())

	// ex #62 Circle, Square, Area, Shape
	ci := circle{2.5}
	sq := square{3, 4}
	fmt.Printf("#62 Area of circle: ")
	info(ci)
	fmt.Printf(", Area of Square: ")
	info(sq)
	fmt.Println("")

	// ex #63
	fmt.Println("#63 Add(5,5):",Add(5,5)) 
}

// ex #63 Test Add
func Add(x int, y int) int {
	return x+y
}

// #ex #62, Square, Circle, Area, Shape

type circle struct {
	radius float64
}

type square struct {
	width  float64
	height float64
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func (s square) area() float64 {
	return s.width * s.height
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("%v", s.area())
}

// ex #61 Create a struct and attach a method to it and use it
type Person struct {
	first string
	age   int
}

func (p Person) Speak() string {
	return fmt.Sprintf("My name is %v and I am %v years old", p.first, p.age)
}

// # ex 58 Call func foo return int, and bar return int, string
func foo() int {
	return 1
}

func bar() (int, string) {

	return 2, "hello"
}

// #ex 59 use a variant input and unfold, and a slice input and range values and return a result
func foovar(vi ...int) int {
	i := 0
	for _, j := range vi {
		i += j
	}
	return i
}

func barvar(vi []int) int {

	i := 0
	for _, x := range vi {
		i += x
	}
	return i
}

