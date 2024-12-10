package main

import "fmt"

// Shape interface requires a Render method
type Shape interface {
	Render() string
}

// Circle struct with a Render method and Resize method
type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

// Square struct with a Render method
type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

// ColoredShape struct that wraps another Shape and adds color
type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

// TransparentShape struct that wraps another Shape and adds transparency
type TransparentShape struct {
	Shape        Shape
	Transparency float32 // a percentage between 0.0 and 1.0
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency*100.0)
}

func main() {
	// Create a Circle with a radius of 2
	circle := &Circle{Radius: 2}
	fmt.Println(circle.Render())

	// Create a ColoredShape based on the Circle with color "Red"
	redCircle := ColoredShape{Shape: circle, Color: "Red"}
	fmt.Println(redCircle.Render())

	// Create a TransparentShape based on the ColoredShape with 50% transparency
	rhsCircle := TransparentShape{Shape: &redCircle, Transparency: 0.5}
	fmt.Println(rhsCircle.Render())
}
