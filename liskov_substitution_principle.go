package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	height, width, length int
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func (r *Rectangle) GetWidth() int {
	return r.height
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

type Square struct {
}

func NewRectangle(height, width, length int) *Rectangle {
	return &Rectangle{height, width, length}
}

func NewSquare(dimension int) *Rectangle {
	return &Rectangle{dimension, dimension, dimension}
}

func UseIt(size Sized) {
	actualArea := size.GetHeight() * size.GetWidth()
	fmt.Printf("Area:%d", actualArea)
}

func testShapes() {
	r := NewRectangle(1, 2, 3)
	s := NewSquare(1)
	fmt.Print(r, s)
}
