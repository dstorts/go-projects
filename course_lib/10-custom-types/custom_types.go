package main

import (
	"fmt"
	"unsafe"
)

// Define a custom type 'special_int' based on int32
// where below, we will add methods to this custom type
type special_int int32

func (si *special_int) memory_info() {
	fmt.Printf("Memory Address: %p\n", si)
	fmt.Printf("Memory Used   : %d bytes\n", unsafe.Sizeof(si))
}

func (si special_int) is_even() bool {
	return si%2 == 0
}

func main() {
	var mySpecialInt special_int = 42

	fmt.Println("Raw Prints ----")
	fmt.Println("Address of mySpecialInt:", &mySpecialInt)
	fmt.Println("Method Based Info ----")
	fmt.Println("Value of mySpecialInt:", mySpecialInt)
	mySpecialInt.memory_info()
	fmt.Println("Is mySpecialInt even?", mySpecialInt.is_even())
}
