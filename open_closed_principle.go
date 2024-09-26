package main

import "fmt"

const (
	small Size = iota
	medium
	large
)

const (
	red Color = iota
	green
	yellow
)

type Size int
type Color int

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

type Specification interface {
	IsSatisfied(*Product) bool
}

func (c Color) IsSatisfied(p *Product) bool {
	return p.color == c
}

func (s Size) IsSatisfied(p *Product) bool {
	return p.size == s
}

func (f *Filter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", green, large}

	products := []Product{apple, tree, house}

	f := Filter{}

	for _, v := range f.Filter(products, green) {
		fmt.Println(v)
	}
	for _, v := range f.Filter(products, large) {
		fmt.Println(v)
	}
}
