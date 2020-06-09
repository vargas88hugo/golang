package main

import "fmt"

// Point strut that represents a point
type Point struct {
	X, Y int
}

// Circle struct that represents a circle
type Circle struct {
	Point
	Radius int
}

// Wheel struct that represents a wheel
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}

	fmt.Printf("%#v\n", w)
	// main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}

	w.X = 42

	fmt.Printf("%v\n", w)
	// {{{42 8} 5} 20}
}
