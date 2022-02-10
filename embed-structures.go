package main

import "fmt"

type coords struct {
	x, y int
}

type circle struct {
	radius int
	coords coords
}

func main() {
	var ring circle
	ring.radius = 15
	ring.coords.x = ring.radius
	ring.coords.y = ring.radius

	fmt.Printf("Diameter :%v\n", ring.getDiameter())
	fmt.Printf("Point X:%v Y:%v\n", ring.coords.x, ring.coords.y)
}

func (c circle) getDiameter() int {
	return 2 * c.radius
}
