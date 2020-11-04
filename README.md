# design-pattens-go

## Principles (Robert C. Martin)
1. Single Responsibility Principle (SRP)
2. Open-closed Principle (OCP)
3. Liskov Substitution Principle (LSP)
4. Interface Segregation Principle (ISP)
5. Dependency Inversion Principle (DIP)


### SRP
Sates that a type should have one primarily responsibility.
Let's see how you can adhere to both the SRP and as well as how you can break.
Let's imagive you are making a journal. you are going to record your thoughts.

=> we can simply break this single responsibility by adding functions which
deals with another concerns
* Seperation of Concerns => different problems(/concerns) which the system solves have to
* reside in a different construct, whether attached to different structures or
* residing in different packages. But basically you need to seperate that.
  Because that is an anti-pattern is called a GOD object.
* A God object - where you take everything and all functionality and just add
  that in a single package - BAD IDEA.

```
package main

var entryCount = 0
type Journal struct {
	enteries []string
}

// primarily resp of the journal is to add an entry and return the index
func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.enteries = append(j.enteries, entry)
	return entryCount
}

// removing an entry from a journal for a particular index i
func (j *Journal) RemoveEntry(index int) {
	//...
}

// Now let's say you did that functionality till above to add and remove an
// entry and then now you decided you need persistence for this journal too.
// So you added `Save`
func (j *Journal) Save(filename string) {
	_ := ioutil.WriteFile(filename, []byte(j.String(), 0644)
}

// Now let's say you decided to make a string for the journal entry.
func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) Load(filename string) {}

func (j *Journal) LoadFromWeb(url *url.URL) {}

func main() {

}
```
* So what we're doing here is we're breaking the single responsibility principle because the responsibility
* of the journal is to do with the management of the entries the responsibility of the journal is not
* to deal with persistence persistence can be handled by a separate component whether it's a separate
* package or whether for example you want to have an actual struct that has some methods related to persistence.
* But the point is you don't necessarily want to have persistence as part of the journo as methods on
* the journal and you might be wondering well why not.

* This is because there are other parts also in the system which also need loading and saving. So we want to keep
* that somewhere. Now both that and this persistence of journal have this in common. So we should keep that
* in a seperate package for that or a seperate struct.

##### seperate package example.
```
// so if you see this func can exist in a different package for just doing the persistence.
// this will be used to save not just journal and also for other objects.
var LineSeperator = "\n"
// lineseperator can be different for windows and linux
// and if we have persistence seperated out in seperate package then managing all of this will be more easier.
func SaveToFile(j *Journal, filename string) {
	_ = ioutil.Writefile(filename, []byte(strings.Join(j,entries, LineSeperator)), 0644)
}
```
#### seperate struct example for persistence.
```
type Persistence struct {
	lineSeperator string
}

// this will be a method rather than a function.
func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.Writefile(filename, []byte(strings.Join(j,entries, LineSeperator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I laughed a lot today")
	j.AddEntry("I learned go design patterns today")
	fmt.Println(j.string())

	// now to save the journal.
	// 1. we can either use a package and call the fucntion from it.
	SaveToFile(&j, "journal.txt")

	// 2. or we can create a seperate persistence struct.
	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}

```

##### Recap:
1. Journal has a SRP. storing, removing and manipulating those entries.
   this means it has a single responsibility.
2. But when it comes to some other functionality like persistence. Since
   this concern can be used for other things like how books or how other structs
   are saved. So we should seperate these two outs.
   So it will be much better with lesser code duplication instead of doing same
   things for all different types of struct e.g. journal, books, manuscripts.
   We can implement persistence part seperately to avoid this.

## Open-Closed Principle
1. basically states that types should be open for extension but closed for modification.
2. we are also going to talk about enterprise pattern called specification.

* Let's imagine we are maintaining an online stores where we are selling widgets.
* Now let's say our customers wants to filter the items by some criteria.

```
package main

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

type Product struct {
	name string
	color Color
	size int
}
```

3. So our 1st requirement is to create some filter type so that we can filter it by color.
```
type Filter struct {
	//
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	// gather and match which matches the color and return those into return array.

	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &product[i])
		}
	}
	return result
}
```

```
func main() {
	apple := Product{"Apple, green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	// old methods
	fmt.Printf("Green products (old):\n")

	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

}

// Output
- Apple is green
- Tree is green
```

* Now let's say the specification changes that now we also have to add a filter by size
* So we will one more function name saying `FilterBySize`
```
func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	// gather and match which matches the color and return those into return array.

	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &product[i])
		}
	}
	return result
}
```

* Now again the specification changed and they asked to make both filter by size and filter by color.
* again doing copy paste and modify for both `FilterBySizeAndColor`
```
func (f *Filter) FilterByColorAndSize(products []Product, color Color, size Size) []*Product {
	// gather and match which matches the color and return those into return array.

	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color && v.size == size {
			result = append(result, &product[i])
		}
	}
	return result
}
```

* => So this shows violation of open close principle. Since we are going and modyfing and inerfering with
* what is being written.
* So what open close principle says is that we should be able to open to extent the functionality but not
* by modyfing existing tested functions and functionality but by extending it.

* -> so we want to leave Filter type alone and we should not keep adding more methods to it.
* we should have some extendible setup.
* So we should then use specification pattern. It has bunch of interfaces for adding functionality.
* So 1st thing we do is we make "Specification iterface"

```
type Specification interface {
	IsSatisfied(p *Product) bool
}
```

* we are testing whetehr or not the product mention satisfied the given criterias.
* So then we specifiy different creterias needed to satisfy the given specification.

e.g.
```
type ColorSpecification struct {
	color Color
}

//and then we will have a method defined on color spec.
func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}
```

* Similarly we can have size specification and things like that.

```
type SizeSpec struct {
	size Size
}

func (s SizeSpec) IsSatisfied(p *Product) bool {
	return p.size == s.size
}
```


* betterfilter now, we are going to buiild a diff. filter. now we will have
* a betterfilter. this is something we will never modify ???
* so basically the idea is that given a bf so you also specify a method on it.
```
type BetterFilter struct {}

func (f *BetterFilter) Filter(products []Product, spec Speficiation) []*Product {
	// now we have the spefication object given to us.

	result := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
}
```

* Now above is the setup that we have and now let's see how to use it.
* So both old and new give us same results. but below will be giving us more flexibility.
* 2nd approach with spec is more flexible since if we want to filter by a type, we just need
* to add define a new spec by making a size spec and then use it.
* this follows Open Close principle

* So the types in this case. The interface is open for extension but it is closed for modication.
* Means we will not modify this interface. and we will not modify the betterfilter.
* So at no point I am going to modify the struct that I already created.

```
func main() {
	apple := Product{"Apple, green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	// old methods
	fmt.Printf("Green products (old):\n")

	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Println("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" = %s is green\n", v.name)
	}
}
```

* for size and color we just need to add a compositeSpecification.
* which is just a combination.

```
type AndSpecification struct {
	first, second Specification
}

func (a Andspecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) &&
		a.second.IsSatisfied(p)
}
```

```
func main() {
	apple := Product{"Apple, green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	// old methods
	fmt.Printf("Green products (old):\n")

	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Println("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" = %s is green\n", v.name)
	}

	// composite way of doing this.
	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}

	fmt.Printf("Large green products:\n")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
```

#### Recap/summary
1. types are open for extension. grab and interface and implement somewhere in your code.
but they are closed for modification. So we should not break the existing spec by modifying it.
but instead we should extend it over and over again but by not adding methods inside that type interface.


## Liskov Substitution Principle
1. not primarily applicable to go. since inheritence is not there in go.
2. states that if you have some APIs which takes a base class and if it correctly works with base class then it should
   also work correctly with derived class.
3. But since we don't have inheritance in go. This is not applicable to go.
4. But let's try this by doing some example with some variation.

* you are trying to deal with geometric shapes of rectangular nature. And you decide to have an interface called `Sized`
* that allows you to specify some methods on those constructs e.g. `getters` and `setters`
* So now you suppose you have a type `Rectangle`
```
type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}
```

* Now we can go and implement the `Sized` interface on type struct `Rectangle`

```
func (r *Rectagle) GetWidth() int {
	return r.width
}

func (r *Rectagle) SetWidth(width int) int {
	r.width = width
}
func (r *Rectagle) GetHeight() int {
	return r.height
}
func (r *Rectagle) SetHeight(height int) {
	r.height = height
}
```

* Now we can write a function `UseIt` which will use this `Sized` example
* and UseIt function relies upon the `Sized` interface strictly.
* So any implementers of the `Sized` types should not break the core principle of `Sized` interface

```
func UseIt(size Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)

	// expected area
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()

	fmt.Println("Expected an area of", expectedArea, "but got an area of",
			actualArea)
}
```

* we can try donig that as an example.
```
func main() {
	rc = &Rectangle{2, 3}
	UseIt(rd)
	// o/p we get both area as 20
}
```

* until you try to break the Liskov principle everything is working.

```
type Square struct {
	Rectangle  // by just spcifiying this we uplift rectangle.width and height
}

// but let's say you enforce that width will always be equal to height  => this here is the problem
func NewSquare(size int) *Square {
	sq := Square {}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width // here you break the Liskov substitution principle
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height // here you break the Liskov principle
}

// now we took a rectangle and we also thought it will work with square.
// but this does not work.

Now when we call use it on square this will break.
```

`UseIt` will work on Rectangle object which is implenting the same interface type `Sized`
but the same function `UseIt` is now breaking on `Square` object which actually happens
to take the same Rectangle type object and implements the same interface

This is why -

* We expected an area of 50 and we got an area of one hundred. So what happened here.
* Well obviously what happened is this call to set height actually set not just the height. It also set the width.
* So the internal width of the square became inconsistent with the value of this variable right here and
* as a result we're getting different values for expected area and actual area.

* So our Liskov substition principal basically states that if you continue to use generalizations
* like interfaces for example then you should not have inheritance or you should not have implementers
* of those generalizations break some of the assumptions which are set up at the higher level.
* So at the higher level we kind of assumed that if you have a sized object and you set its height you
* are just setting its height not both the height and the width and here what happened is we broke this
* assumption by setting both the width and the height which is a noble goal.
*
* You can see how somebody would try to enforce the square invariant by setting both within the height.
* It's a noble goal but it doesn't work. And it actually violates the Liskov substitution principle.

* So the Liskov substitution principle is actually one of those situations where there is no right
* answer there is no right solution to this problem.

* => so here is how we could do for this square.
* the behavior of a particular implementer of a particular type like in this case `Sized` interface
* should not break the core fundamental behavior that we reply on.
```
type Square2 struct {
	size int // width, height
}

func (s *Square) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}
```


## Interface Segregation Principle (mostly simple)
1. states that you shouldn't put too much into the one single interface.
2. It makes more sense to break that interface into several smaller interface.


* e.g. we have a `Document` type struct with some info about the document
* we want to make the interface which allows ppl to build diff. machines/ diff. constructs operating on the documents.
* e.g. printing, faxing, scanning the document.

3. So one approach is to make one single interface. in which we will have all the methods define for
* for printing the document, faxing the document and scanning it.
* So this is generally ok, if we are looking for something like multi-function printer.
* e.g. my printer can do both printing and scanning.

```
type Document struct {}
type Machine interface {
	Print (d Document)
	Fax (d Document)
	Scan (d Document)
}

type MultiFunctionPrinter struct {

}
// now you go and implement all the interface methods.
func (m MultiFunctionPrinter) Print(d Document) {}
func (m MultiFunctionPrinter) Scan(d Document) {}
func (m MultiFunctionPrinter) Fax(d Document) {}



```

4. However imagine a diff. situation where someone is weorking with old fashioned printer. which does not have scanning and faxing functionality.
* but bcoz we want to implement this. say some machine reply on this. Because `Machine` interface will force us to implent this.
* So we will have to implement all methods which are for multi-function device.
* But we don't know what to do with Faxing and Scanning. So we will leave a panic method there.
```
type OldFashionedPrinter struct {

}

func (o OldFashionedPrinter) Print(d Document) {
}

// Deprecated - : (not actually since the method is not deprecated. we are lying to user.
func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}
func (o OldFashionedPrinter) Scan(d Document) {}

func main() {
	ofp := OldFashionedPrinter{}
	ofp.Scan() // some ides will complain that this method is deprecated.
	// But what has really happened is we have created a problem by adding lot of functionality in a single interface.
	// and then we are asking everyone to implement this interface by forcing them.
}
```

* interface segregation principle.
* try to break the interface into seperate part so that folks can implement those individually since those may not be needed by a single interface.

e.g.
```
type Printer interface {
	Print (d Document)
}
type Scanner interface {
	Scan (d Document)
}

// only a printer.
type MyPrinter struct {
}

// since this does not fo anything except printing so we only implement a Printer interface
func (m MyPrinter) Print(d Document) { }


// here is an example where we can implment both functions by allowing below struct implement
// both the mthods of Print and Scan
type Photocopier struct {}

func (p Photocopier) Print(d Document) {
	panic("implement me")
}
func (p Photocopier) Scan(d Document) {
	panic("implement me")
}
```

* we can also create a multi-functional interface like this by adding diff. interfaces into a composite interfaces
* combining interfaces is fairly easy.

```
type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax
}
// you want to build a multi-function machine.
// you already got a printer and scanner implemented as a seprate component.
// we can use a decorator design pattern.

//decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	// here we will re-use functionality of printer.
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	// here we will re-use functionality of printer.
	m.printer.Scan(d)
}

// you can see with segration approach you have more granular approach.
```

#### Recap
1. We should segregate interfaces into multiple smaller interfaces since we anyway can always combine and create a composite interface later if we need it.


## Depeendency Inversion Principle
1. does not have anything to do with dependency injection. they should be kept seperate.
2. We are only talking about the principle and not the actual mechanism.
3. States that high level module should not depend on low level modules and both of them should depend upon abstraction.

* let's imaging you are doing some research where you are trying to model relationships between different ppl.
*
```
type Relationship int

// one person can be parent of another person.
// or you could be child etc...
const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// other type of info like dob etc.
}
```

* So now how do we model relationship between different ppl.

```
type Info struct {
	from *Person
	relationship Relationship
	to *Person
}
```

* so you can for e.g. say that "John" is the "parent" of "Jill"
* Now how will we store this information.

```
type Relationships struct {
	relations []Info
}
```
* so what do we want to get from all of this modeling. maybe we need to get all child of one person.
* But the 1st thing we need to do is we need to be able to add those relationships.

```
// so you specify the parent and the child and they are both Persons pointer.
func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations,
		Info {parent, Parent, child})
	r.relations = append(r.relations,
		Info {child, Child, parent})
}
```

* and then what we want to be able to do is to be able to perform research on this data

```
// high level module. this is a high level module.
// and relationships is a low level module. Reason is because this is some form of storage
// this can be in some database or from Web.
// But research is some of high level module.
// so high level module should not depend on low level module.
type Research struct {
	// break DIP.
	relationships Relationships
}

// so this does works by allowing us to do research.
func (r *Research) Investigate() {
	relations := r.relationships.relations
	for _, rel := range relations {
		// find all children of "John"
		if rel.from.name == "John" && rel.relationship == "Parent" {
			fmt.Println("John has a child called", rel.to.name)
		}
	}
}

// this above will work. But it has a problem.
```
* Let's see it in action.

```
func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	// low level part
	relationships = Relationships{}
	relations.AddParentAndChild(&parent, &chidl1)
	relations.AddParentAndChild(&parent, &chidl2)

	// now we can perform research
	// and it is going to rely on low level module.
	r := Research{relationships}
	r.Investigate()

	// op is both the childs gets printed.
}
```

* => but there is a big problem. Since the high level module `Research` is usign a low level module.
* now say the low level module decides to change the storage mechanism from []slice to database.
* So then the high level module which is dependent on low level module will break.
* Hence this dependency module Principle should be followed to avoid this.
* Also investigate method should be part of low level module directly, since we can optimize based on
* storage mechanism.

* Let's rewrite this with a right solution `both should depend on abstraction`
```
type Research struct {
	// we will avoid the situation where we are exposing the low level interface of relationships.
	// instead will implement below interface.
	browser RelationshipBrowser
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// now what we can do is we can implement above interface and attach the method to Relationships.
func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	// now since here we are in low level module, we can go about how we will access the low level storage.

	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

```
* now all finding of the children is put into low level module and then we can re-write the high level module.
* so basically what I got the idea is that between high level module and low level module there should be some abstractions created
* which the high level module should access.
* basically in golang a struct can have a method inside it by embedding an interface or attaching a method to it.
* But if we are attaching a method that is something that we will implement. But a low level module can implement a method
* which then we will call it as a interface which then we can use it in our high level module.

* so basically if the high level module wants to use a function instead of adding relationships struct into the high level
* we just create an interface which implements the low level module and then embed this interface into our high level struct.

```
func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}
```

* this way low level implementation is not exposed to high level. We are just using low level abstraction method which will be created by low level
* for use by high level as an interface. This way high level does not have to use the low level structures/ or definitions.


### Recap -
* high level modules should not depend upon low level modules. low level meaning - hw, storage, communications etc. High level meaning business logic stuff.
* both of them should be seperated by abstractions. imn golang this means interfaces. Other lang. means abstrace and base classes.
* In golang means interfaces.
* This way we are protecting against changes between low level and high level. e.g. if we change the slice to database. then we will only be modifying methods
* of `Relationships` and not the methods of `Research`. Because both should are seperated out by this design principle.


### SRP
1. a type should have only one primary responsibility.
2. Seperation of concerns - different areas of responsibility then we should be putting them in diff. types/packages.

### OCP (open-closed principle)
1. types should be open for extension but closed for modification.
2. if you have written and tested your code then shouldn't jump back into the code to modify.

### Liskov Substition Principle
1. If you have sometype which kind of aggregates another type and thereby acquire all of it's method and so on.
2. you should be able to substitute an embedded type in place of it's embedded part.

### Interface Segregation Principle (the most important one)
1. Don't put too much into an interface, split into seperate interfaces.
2. YAGNI - you ain't going to need it.

### Dependency Inversion Principle
1. high level modules should not depend upon low level ones. But should use abstractions.
2. Here you always segregate between low level and high level. Define data structures and methods which operate on low level.
* then you specify implement a abstraction method by the low level module which can be used by a high level module.
* In the high level module you define an interface which just uses this low level abstraction method exported.
* Then in the high level module you define a struct where we put this collection of interfaces which we will need to call.
* So the high level module will just call `r.browser.FindAllChildrenOf` - just refer previous e.g.
* (This as per my understanding also will inherently cover open-closed principle (and all others) as well.
* Since we have segregated most of the different functionality related modules into seperate packages)

### The above are 5 solid design principles which we will need while discussing the different design patterns.
