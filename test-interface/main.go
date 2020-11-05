package main

import (
	"fmt"
	"github.com/riteshharjani/design-pattens-go/test-interface/lowlevel"
)

// this shows how to segregate low level and high level routines.
type highlevel_struct struct {
	high_interface lowlevel.ExposedInterface
	// we can add more stuff here
}

func main() {
	var st highlevel_struct
	st.high_interface = lowlevel.Get_ExposedInterface()
	fmt.Println("hello world")

	st.high_interface.Dummy_lowlevel_add_val(1)
	st.high_interface.Dummy_lowlevel_add_val(2)
	st.high_interface.Dummy_lowlevel_print_val()
}
