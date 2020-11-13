package main

import (
	"fmt"
	"strings"
)

// we are going to look into something which is already built into go
// stringbuilder.
// idea is that we need to build up strings of HTML from text elements.
// so from text we need to convert to paragraph.
func main() {
	// e.g. piece of text hello
	hello := "Hello"

	// to make it para from above string we will use stringbuilder
	// this is part of the go SDK.
	// help cat strings together by allowing to write several string one
	// after another in buffer and then gives us a cat result.
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Println(sb.String())
	// o/p
	// <p>Hello</p>

	// now let's say you have more words.
	words := []string{"hello", "world"}
	// reset the sb
	sb.Reset()

	// <ul><li>...</li><li>...</li> <...> </ul>
	// so how do you above doing the string builder
	// below will be one of the way
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")

	fmt.Println(sb.String())
	// o/p:
	// <ul><li>hello</li><li>world</li></ul>

	// but this is not a very effective method.
	// since we know we need a opening and a closing tag
	// in every word.
	// so why not improve this.
	// so why not put something into struct which provide more flexibility
	//
	// This is the motivation of builder.
	// We have some object which we want to build up into some steps.
	// So like above there are lot of steps for building the object.
	// So we need to help client build a HTML page.

	// So instead of adding/typing opening closing tag everytime,
	// we can represent the entire HTML construct as a tree and then learn how to print that tree
	// and give the user a builder component, where they can add an element to thsi tree.
	// w/o knowing where this tree is actually there.

	// e.g.
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Println(b.String())

	// if you now see the output will be like below.
	// o/p :=
	// <ul>
	//   <li>
	//     hello
	//   </li>
	//   <li>
	//     world
	//   </li>
	// </ul>

	// Fluent method is to chain the calls together
	b = NewHtmlBuilder("ul")
	b.
		AddChildFluent("li", "hello").
		AddChildFluent("li", "world")
	fmt.Println(b.String())

}

// now we need to do couple of things. We need these elements to be printable.
// and we also need some indentation
const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

// String() will be calling string() with indentation with size
func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))

	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName: rootName,
		root: HtmlElement{
			name:     rootName,
			text:     "",
			elements: []HtmlElement{},
		},
	}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
}

// if you return the builder then with that we can chain the calls together
func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

/*
func NewHtmlBuilder(rootName string) *HtmlBuilder {

}
*/
