package main

import "fmt"

// by using a functional programming approach.
// this is another way of doing the builder pattern
type Person struct {
	name, position string
}

// a list of modification which areggoing to apply to PErson

type personMod func(*Person)

// inside this we will have list of actions
type PersonBuilder struct {
	actions []personMod
}

func (b *PersonBuilder) Called2(name string) *PersonBuilder {
	b.actions = append(b.actions,
		func(p *Person) {
			p.name = name
		})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func main() {
	fmt.Println("vim-go")
}
