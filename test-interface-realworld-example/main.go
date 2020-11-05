package main

import (
	"fmt"
	"github.com/riteshharjani/design-pattens-go/test-interface-realworld-example/lowlevel_device"
	ld "github.com/riteshharjani/design-pattens-go/test-interface-realworld-example/lowlevel_device_2"
)

func main() {
	fmt.Println("lowlevel_device test")
	device := lowlevel_device.NewDevice()
	device.AddData(1)
	device.AddData(2)
	device.AddData("hello world")
	device.PrintData()

	fmt.Println("Device length: ", lowlevel_device.Len(device))
	// output -
	// Real world test
	// 0: 1
	// 1: 2
	// 2: hello world
	//
	//

	fmt.Printf("\n\nlowlevel_device_2 test:\n")
	ld.AddData_2(1)
	ld.AddData_2(2)
	ld.AddData_2(3)
	ld.AddData_2("hello")
	ld.PrintData_2()
}
