package main

import "fmt"
import "encoding/json"

type Point struct {
	x int
	y int
}

type Square struct {
	X int
}

func main() {
	var p Point
	p.x = 1
	p.y = 1
	// Create
	p = *new(Point) // zero type values
	p.x = 2
	fmt.Println(p)
	p = Point{x:1, y:1}
	fmt.Println(p)

	// Struct field visibility
	s1 := Square{2}
	p1 := Point{2, 2}
	s1JSON, _ := json.Marshal(s1) // Square field are seen by json package
	p1JSON, _ := json.Marshal(p1) // Square field are not visible by json package
	fmt.Println(string(s1JSON)) // returns s1 json
	fmt.Println(string(p1JSON)) // returns {}
	fmt.Println(p1) // returns {2, 2}. Point private fields are visible in same pkg
}