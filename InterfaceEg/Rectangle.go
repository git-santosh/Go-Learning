package main

import "fmt"

type Rectangle struct {
	width, height float64
}

type Shape interface {
	area() float64
}

func getArea(shape Shape) float64 {
	return shape.area()
}
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}
func main() {
	rectangle := Rectangle{width: 10, height: 5}
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))

}
