package main

import "fmt"

/*
Summary:
types are open for extension. grab and interface and implement somewhere in your code.
But they are closed for modification. So we should not break the existing spec by modifying it.
But instead we should extend it over and over again but by not adding methods inside that type interface
*/

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

// Specification is an interface
type Specification interface {
	IsSatisfied(p *Product) bool
}

// till above can be a part of product package

// colorSpecification - we can define below in our package
type colorSpecification struct {
	color Color
}

func (c colorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type sizeSpecification struct {
	size Size
}

func (s sizeSpecification) IsSatisfied(p *Product) bool {
	return s.size == p.size
}

type betterFilter struct{}

func (f *betterFilter) Filter(products []Product, spec Specification) []*Product {

	res := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			res = append(res, &products[i])
		}
	}
	return res
}

type andSpecification struct {
	first, second Specification
}

func (a andSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	bf := betterFilter{}
	fmt.Println("Green products (new):")

	greenspec := colorSpecification{green}
	for _, v := range bf.Filter(products, greenspec) {
		fmt.Printf(" - %s is green\n", v.name)
	}
	//Green products (new):
	//	- Apple is green
	//	- Tree is green

	// composite way of addding two constraints
	largespec := sizeSpecification{large}
	lgspec := andSpecification{greenspec, largespec}

	fmt.Printf("Large green products:\n")
	for _, v := range bf.Filter(products, lgspec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
	// Large green products:
	//   - Tree is large and green
}
