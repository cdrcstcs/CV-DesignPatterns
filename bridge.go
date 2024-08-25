// Here's the equivalent example of the Bridge design pattern in Go:

// Copy
package main

import "fmt"

// DrawingAPI interface
type DrawingAPI interface {
	DrawCircle(x, y, radius float64)
	DrawRectangle(x, y, width, height float64)
}

// DrawingAPI1 struct
type DrawingAPI1 struct{}

func (api *DrawingAPI1) DrawCircle(x, y, radius float64) {
	fmt.Printf("API1 - Drawing Circle at (%.2f, %.2f) with radius %.2f\n", x, y, radius)
}

func (api *DrawingAPI1) DrawRectangle(x, y, width, height float64) {
	fmt.Printf("API1 - Drawing Rectangle at (%.2f, %.2f) with width %.2f and height %.2f\n", x, y, width, height)
}

// DrawingAPI2 struct
type DrawingAPI2 struct{}

func (api *DrawingAPI2) DrawCircle(x, y, radius float64) {
	fmt.Printf("API2 - Drawing Circle at (%.2f, %.2f) with radius %.2f\n", x, y, radius)
}

func (api *DrawingAPI2) DrawRectangle(x, y, width, height float64) {
	fmt.Printf("API2 - Drawing Rectangle at (%.2f, %.2f) with width %.2f and height %.2f\n", x, y, width, height)
}

// Shape interface
type Shape interface {
	Draw()
}

// Circle struct
type Circle struct {
	x, y, radius float64
	api          DrawingAPI
}

func (c *Circle) Draw() {
	c.api.DrawCircle(c.x, c.y, c.radius)
}

// Rectangle struct
type Rectangle struct {
	x, y, width, height float64
	api                 DrawingAPI
}

func (r *Rectangle) Draw() {
	r.api.DrawRectangle(r.x, r.y, r.width, r.height)
}

func main() {
	api1 := &DrawingAPI1{}
	api2 := &DrawingAPI2{}

	circle1 := &Circle{x: 1, y: 2, radius: 3, api: api1}
	circle2 := &Circle{x: 5, y: 7, radius: 11, api: api2}

	rectangle1 := &Rectangle{x: 2, y: 3, width: 6, height: 9, api: api1}
	rectangle2 := &Rectangle{x: 4, y: 6, width: 8, height: 12, api: api2}

	circle1.Draw()
	circle2.Draw()
	rectangle1.Draw()
	rectangle2.Draw()
}
// In this Go implementation:

// DrawingAPI is an interface defining the drawing operations.
// DrawingAPI1 and DrawingAPI2 are concrete implementations of DrawingAPI.
// Shape is an interface representing the abstraction (Shape).
// Circle and Rectangle are implementations of Shape, utilizing a specific DrawingAPI.
// The client code demonstrates how different shapes can use different drawing APIs.
// This Go version mirrors the structure of the Java example while adhering to Go's syntax and idioms. The Bridge pattern still allows for decoupling between abstraction and implementation, facilitating flexibility and extensibility.