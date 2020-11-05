package main

import "fmt"

// Color type
type Color int

// Size type
type Size int

const (
	red Color = iota
	green
	blue
)
const (
	small Size = iota
	medium
	large
)

// Product info
type Product struct {
	name  string
	color Color
	size  Size
}

//Filter struct
type Filter struct{}

// FilterByColor func
func (f *Filter) FilterByColor(products []Product, color Color) []*Product {

	res := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			res = append(res, &products[i])
		}
	}
	return res
}

// FilterBySize - add another specification to filter by size
func (f *Filter) FilterBySize(products []Product, size Size) []*Product {

	res := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			res = append(res, &products[i])
		}
	}
	return res
}

//FilterByColorAndSize - filter by both color and size
func (f *Filter) FilterByColorAndSize(products []Product, color Color, size Size) []*Product {
	// gather and match which matches the color and return those into return array.

	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color && v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	// old methods
	fmt.Println("Green products (old):")

	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}
	// output
	//Green products (old):
	// - Apple is green
	// - Tree is green
}
