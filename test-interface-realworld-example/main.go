package main

import (
	"fmt"
	"github.com/riteshharjani/design-pattens-go/test-interface-realworld-example/lowlevel_device"
)

func main() {
	fmt.Println("Real world test")
	device := lowlevel_device.NewDevice()
	device.AddData(1)
	device.AddData(2)
	device.AddData("hello world")
	device.PrintData()

	// output -
	// Real world test
	// 0: 1
	// 1: 2
	// 2: hello world
}
