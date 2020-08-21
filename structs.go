package main

import "fmt"

func main() {
	type Point struct {
		x int
		y int
	}

	var p Point
	p.x = 1
	p.y = 1
	// Create
	p = *new(Point) // zero type values
	p.x = 2
	fmt.Println(p)
	p = Point{x:1, y:1}
	fmt.Println(p)
}