package main

import (
	"fmt"
	"strings"
)

// So one question you might be asking is how do I get the uses of my API to
// actually use my builders as opposed to stop messing with the objects directly.

// And one approach to this is you simply hide the objects that you want
// your users not to touch.
// e.g. say you have an API of some kind for sending emails.
//

type email struct {
	from, to, subject, body string
}

// one of the problems one may have is that we want our email struct to be
// fully specified. So then may have to write a validator whcih could validate
// each email
// or instead we can create a builder so that the user can invoke methods
// on that builder in order to build out that email.
//
// (don't expose) -
// Also note that we can keep the name of the struct email in lowercase
// via which the user cannot modify or access the email struct directly
// since it is not exported from this package.

// we can make EmailBuilder which could be accessible to user.
// but it's not going to expose the different parts of the email directly.

type EmailBuilder struct {
	email email // aggregator
}

// fluent interfaces
func (b *EmailBuilder) From(from string) *EmailBuilder {
	// (we can as well provide some validation here)
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	b.email.from = from
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(sub string) *EmailBuilder {
	b.email.subject = sub
	return b
}

// so how so you do this .. that the user only uses your EmailBuilder

func sendMailImpl(email *email) {
	fmt.Println("Email sent", email)
}

type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}

func main() {
	fmt.Println("builder-params")
	// SendEmail(func (b *EmailBuilder {})
	SendEmail(func(b *EmailBuilder) {
		// below is a initializer routine
		// where the builder pointer passed in SendEmail
		// is used by the client to init the fieds of email.
		b.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello where do you want to meet")
	})
}
