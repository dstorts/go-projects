package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Name string
	Age  int32
}

func main() {
	var a int32 = 42
	var b int16 = 1
	//the following, var c, is a slice declaration of 10 int16 elements
	var c []int32 = make([]int32, 10)
	var d int32 = 27
	var e_str string = "Once upon a time there was a gopher. He was a scared little gopher and sadge."

	fmt.Println("a:", a)
	fmt.Println("a address:", &a)
	fmt.Println("b:", b)
	fmt.Println("b address:", &b)
	fmt.Println("c:", c)
	fmt.Printf("c address: %p\n", c)
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d] address: %p\n", i, &c[i])
	}
	fmt.Println("d:", d)
	fmt.Println("d address:", &d)

	// Calculate and print the size of the struct instance
	fmt.Printf("Size of c[] instance: %d bytes\n", unsafe.Sizeof(c))

	fmt.Printf("Size of c[0] instance: %d bytes\n", unsafe.Sizeof(c[0]))

	fmt.Println("0xc000012504 - 0xc0000124e0 = ", (0xc000012504 - 0xc0000124e0))

	fmt.Println("e_str:", e_str)
	fmt.Println("e_str address:", &e_str)
	fmt.Printf("e_str size in memory: %d bytes", unsafe.Sizeof(e_str))
}
