package main

import "fmt"

// The iterator pattern is all about traversing the data structures.

// Motivation -
// Iterator (traversal) is a core functionality of various data structures
// An iterator is a type that faciliates the traversal.
// 	Keeps a pointer to the current element within a collection and then it has mechanism of advancing.
// 	Knows hot to move to a different/next element.
// Go allows iteration with range
// 	Built in support in many objects (arrays, slices, maps etc)
// 	Can be supported in user-define struct
//

// e.g.
// To iterate meaning let's go over a selection/every element.

type Person struct {
	FirstName, MiddleName, LastName string
}

// what if someone wants to go over every single name in the Person
// one approach is to everything in an array and since array has an built-in
// iterator.
func (p *Person) Names() [3]string {
	return [3]string{p.FirstName, p.MiddleName, p.LastName}
}

// 2nd approach
// Another approach is to use a generator
// using channels and go routines
func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()
	return out
}

// 3rd approach
// a seperate struct. c++ commonly uses it.
// we need to make newtype of struct which will have a pointer to the struct which we want to iterate upon
// and a current value to know where we are while during iteration.
// Pointer is needed to avoid copying the data everytime.
type PersonNameIter struct {
	person  *Person
	current int
}

func NewPersonNameIter(person *Person) *PersonNameIter {
	return &PersonNameIter{person, -1}
}

// move the iterator forward and checks provides a end condition on when
// to break the iteration.
func (p *PersonNameIter) MoveNext() bool {
	p.current++
	return p.current < 3
}

// This will return the value of the current iterator.
func (p *PersonNameIter) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	}
	panic("illegal case")
}

func main() {
	p := Person{"Alexander", "Graham", "Bell"}

	// method 1
	for _, name := range p.Names() {
		fmt.Println(name)
	}

	// method 2 - (generator)
	for name := range p.NamesGenerator() {
		fmt.Println(name)
	}

	// method 3 - iterator (This is what iterator design pattern is all about)
	for it := NewPersonNameIter(&p); it.MoveNext(); {
		fmt.Println(it.Value())
	}

	/* Output of above
	* Alexander
	* Graham
	* Bell
	* Alexander
	* Graham
	* Bell
	* Alexander
	* Graham
	* Bell
	 */

}
