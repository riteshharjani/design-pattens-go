package main

import "fmt"

// in most situations a single builer is sufficient to build a particular obj.
// But there are situations where you need more than one builder way.
// You need to somehow separate the process of building up the different aspects of a particular type.
//
// e.g.
type Person struct {
	// two particular types of info which we want to build up
	StreetAddress  string
	PostCode, City string

	// job info
	CompanyName, Position string
	AnnualIncome          int
}

// So imagine you want to have a separate builders for building up the address
// information and for building up the job information.
// so how should we do it.

// start with PersonBuilder
type PersonBuilder struct {
	person *Person
}

// obviously we have to initialize it.
// so instead of adding everything inside the Personbuilder we can have
// seeprate builders for add and job builder. then we agregate it,
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

// but from PersonBuilder we want to be able to provide interfaces whcih are
// provided by PersonAddressBuilder and PersonJobBuilder.

// we can provide a utility method which gives us that
func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	// PersonAddressBuilder{*b}
	// above is nothing but same as below.
	// not that inside b person is a pointer
	// and the above method assigns person = (*b) which is also a *Person
	return &PersonAddressBuilder{
		PersonBuilder: PersonBuilder{
			person: b.person,
		},
	}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (it *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = streetAddress
	return it
}

func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (it *PersonAddressBuilder) WithPostCode(postcode string) *PersonAddressBuilder {
	it.person.PostCode = postcode
	return it
}

func (pjb *PersonJobBuilder) Earning(income int) *PersonJobBuilder {
	pjb.person.AnnualIncome = income
	return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) At(company string) *PersonJobBuilder {
	pjb.person.CompanyName = company
	return pjb
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().
		At("123 London road").
		In("London").
		WithPostCode("SW12BC").
		Works().
		At("Facebook").
		AsA("Progreammer").
		Earning(123000)
	person := pb.Build()
	fmt.Printf("%p\n", pb.Lives().person)
	fmt.Printf("%p\n", pb.Works().person)
	fmt.Printf("%p\n", pb.person)
	fmt.Println(person)
}
