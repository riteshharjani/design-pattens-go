package main

import "fmt"

// e.g. of mediator design pattern is a simulation of chat room.
// begin by definiting a Participant of chat room "Person"

type Person struct { // this does not know anything about other people
	Name string
	// and it should have a pointer to mediator
	// chat room allows diff people communicate with one another w/o
	// being aware of one other person
	Room    *Chatroom // this is a mediator
	chatLog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

// you should be able to recieve a msg from
// a) another Person
// b) or directly from Chat room about some system msgs
func (p *Person) Receive(sender, message string) {
	// format the msg
	// whose chat session we are actually in
	// and append in the chat log of the person
	s := fmt.Sprintf("%s: %s\n", sender, message)
	fmt.Printf(`[%s's chat session]: %s`, p.Name, s)

	p.chatLog = append(p.chatLog, s)
}

// method from Person to say/chat a msg
func (p *Person) Say(message string) {
	// p.Room is out mediator
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(who, message string) {
	p.Room.Message(p.Name, who, message)
}

// let's define room now.

type Chatroom struct {
	people []*Person // we could have added a map using the key as the name
	// then the search would be O(1)
}

// let's define a way of Broadcasting
func (c *Chatroom) Broadcast(source, message string) {
	for _, p := range c.people {
		if p.Name != source {
			p.Receive(source, message)
		}
	}
}

// let's define a way of messaging one other
func (c *Chatroom) Message(src, dst, msg string) {
	for _, p := range c.people {
		if p.Name == dst {
			p.Receive(src, msg)
		}
	}
}

func (c *Chatroom) Join(p *Person) {
	// let's say when anyone joins then we do a broadcast to everyone
	joinMsg := p.Name + " joins the chat"
	c.Broadcast("Room: ", joinMsg)
	p.Room = c
	c.people = append(c.people, p)
}

func main() {
	room := Chatroom{}
	john := NewPerson("John")
	jane := NewPerson("Jane")

	room.Join(john)
	room.Join(jane)

	john.Say("hi room")
	jane.Say("oh, hey john")

	simon := NewPerson("Simon")
	room.Join(simon)
	simon.Say("Hi everyone")

	jane.PrivateMessage("Simon", "Glad you could join us")
	// o/p of above
	// [John's chat session]: Room: : Jane joins the chat
	// [Jane's chat session]: John: hi room
	// [John's chat session]: Jane: oh, hey john
	// [John's chat session]: Room: : Simon joins the chat
	// [Jane's chat session]: Room: : Simon joins the chat
	// [John's chat session]: Simon: Hi everyone
	// [Jane's chat session]: Simon: Hi everyone
	// [Simon's chat session]: Jane: Glad you could join us
	// ^^^ only simon recieve this since it is a private msg

}
