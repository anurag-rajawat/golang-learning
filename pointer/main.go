package main

import "fmt"

func main() {
	// A Pointer is a variable that is used to store the memory address of another variable.
	// We used & (ampersand) operator to refer to a variable's memory address.
	// and * is used to dereference value
	num := 22
	ptr := &num
	fmt.Printf("Memory address of num is %v\n", ptr)
	fmt.Printf("Value of num is %v\n", *ptr)

	// We can change the value through the pointer but there is nothing like pointer-arithmetic like C/C++
	// We can also create a pointer to pointer
	*ptr = 2003
	fmt.Printf("Value of num is %v\n", *ptr)

	name := "Old Name"
	fmt.Printf("My old name is %v\n", name)
	changeName(&name)
	fmt.Printf("My new name is %v\n", name)

	// 2. New function
	// New fn is used to initialize a pointer and returns a pointer
	// <name> := new(<type>)
	val := new(int)
	fmt.Printf("Memory address of val is %v and its value is %v\n", val, *val)
}

// Pointers as function args
func changeName(name *string) {
	*name = "new name"
}
