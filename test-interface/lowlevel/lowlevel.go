package lowlevel

import "fmt"

type ExposedInterface interface {
	Dummy_lowlevel_print_val()
	Dummy_lowlevel_add_val(int)
}

// so basically the design should be that define low level struct
// and then we can come up with an interface like above which we need to expose
// to high level to call our routine.
type lowlevel_struct struct {
	dummy_val []int
}

// so all of these below functions are exposed to high level layer.
// since that interface is exposed
// But this can operate on structures which are not exposed.
func (l *lowlevel_struct) Dummy_lowlevel_print_val() {
	for _, v := range l.dummy_val {
		fmt.Println("print_from dummy lowlevel_struct", v)
	}
}

func (l *lowlevel_struct) Dummy_lowlevel_add_val(val int) {
	l.dummy_val = append(l.dummy_val, val)
	fmt.Println("add_from dummy lowlevel_struct", val)
}

// this function should be something like NewDevice
func Get_ExposedInterface() ExposedInterface {
	return &lowlevel_struct{}
}
