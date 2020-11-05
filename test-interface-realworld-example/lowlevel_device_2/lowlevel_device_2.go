package lowlevel_device_2

import "fmt"

// I think this it the right way of doing this.
// from lowlevel package we can expose the APIs which will add, print or manipulate the data
// but let's say if we want something from high level to implement
// then we should expose an interface which the high level should implement
// which we will be calling.
// e.g. See sort package
// type Interface interface {} mandates to implement Len, Swap, and Less methods
// on it's structure. If there is any structure which implements such calls this also means
// that it is of type Interface.
// then calling Sort(data []Person) will be interpreted inside the library package as
// Sort(data Interface).
// And then the package will be free to call those methods attached to Person struct
// or say Interface type for low level lib
type lowlevel_device struct {
	storage []interface{}
}

var ld lowlevel_device

func AddData_2(val interface{}) {
	if ld.storage == nil {
		ld.storage = make([]interface{}, 0)
	}
	ld.storage = append(ld.storage, val)
}

func PrintData_2() {
	for i, v := range ld.storage {
		fmt.Printf("%d: %v\n", i, v)
	}
}
