package lowlevel_device

import "fmt"

// we can add several other functions which we feel should be exposed to high level.
// this is what the high level will be calling.
type Device interface {
	AddData(interface{})
	PrintData()
}

type lowlevel_device struct {
	storage []interface{}
}

func (d *lowlevel_device) AddData(v interface{}) {
	d.storage = append(d.storage, v)
}

func (d lowlevel_device) PrintData() {
	for i, v := range d.storage {
		fmt.Printf("%d: %v\n", i, v)
	}
}

func Len(d Device) int {
	return len(d.(*lowlevel_device).storage)
}

func NewDevice() Device {
	return &lowlevel_device{}
}

// FURTHER NOTE:-
// Interface is just a way of exposing of what is expected out of an API
