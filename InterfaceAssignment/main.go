package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct{
	height float64
	base float64
}

type square struct{
	sideLength float64
}

func main() {
	newSquare := square{45}
	newTriangle := triangle{4,12}

	
	printArea(newTriangle)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (triangle) getArea(height float64, base float64) float64 {
	return 0.5*base*height
}

func (square) getArea(sideLength float64) float64 {
	return sideLength*sideLength
}