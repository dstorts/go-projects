package main

import "fmt"

func main() {

	age := 32
	age_pointer := &age // '&' stores the memory address of age variable
	/*
		You could define a pure pointer variable like this:

			var pure_pointer *int

		Notice: - the *<datatype> declares a pointer variable
				- the default value of a pointer is nil (it points to nothing, but is the reserved nil value)
		RULES:
			1. You cannot do arithmetic with pointers (like in C/C++)
			2. You cannot assign an integer value to a pointer variable
			3. You can only assign the address of a variable of the same datatype to a pointer variable
	*/
	Get_Adult_Years(&age)
	fmt.Println("Adult Years Equals = ", age)
	fmt.Println("Age Pointer: ", age_pointer)       // because age_pointer was assigned the address of age, it value is actually a memory address
	fmt.Println("Value at Address: ", *age_pointer) // '*' dereferences the pointer memory address to get the value at that memory address
}

func Get_Adult_Years(age *int) {
	// a pointer type variable is always just and address.
	// using a pointer type argument means that we do NOT create
	// a new variable in memory for this function every time it is called.
	// so 'age *int' is always an address of an int passed to this function. In this case from the scope of main()
	// and manipulates that value directly.

	// return *age - 18
	*age = *age - 18
	//above:
	// Right Hand Side: dereference the passed memory address to get the value and manipulate it
	// Left Hand Side: reference the passed memory address and store the new value at that address
	// NO need for a return value because we just modified the original variable in main()
}
