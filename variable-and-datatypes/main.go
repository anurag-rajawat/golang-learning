package main

import "fmt"

func main() {
	// Go is very strict so whatever we declared we need to use it otherwise our program will not compile
	// If you don't want to use it you can simply replace that variable with an underscore ( _ )

	// Any variable declared without an explicit initial value are given their zero value
	// int and float are assigned as 0, bool as false, and string as an empty string.
	// For other datatypes the zero value is 'nil'

	// 1. Declaration without initialization
	// var <name> <data_type>
	// Annotation verb sheet to format arguments -> https://pkg.go.dev/fmt
	var usr string
	fmt.Printf("The value of 'usr' without initialization is: %q\n", usr)

	// 2. Declaration with initialization
	var user string = "Anurag"
	fmt.Println("'user' with declaration and initialization:", user)

	// We can omit datatype because go compiler will infer it for us
	var user2 = "I'm user2"
	fmt.Println("'user2' is", user2)

	// 3. Multiple declaration
	var foo, bar, baz string = "foo", "bar", "baz"
	fmt.Printf("Multiple declaration of foo, bar and baz and their corresponding values are: %s, %s and %s\n", foo, bar, baz)
	// As said above we can omit datatype
	var foo2, bar2, baz2 = "foo2", "bar2", "baz2"
	fmt.Printf("Multiple declaration of foo, bar and baz and their corresponding values are: %s, %s and %s\n", foo2, bar2, baz2)
	// OR
	var (
		foo3 string = "foo"
		bar3 string = "bar"
		baz3 string = "baz"
	)
	fmt.Printf("Multiple declaration of foo, bar and baz inside a block and their corresponding values are: %s, %s and %s\n", foo3, bar3, baz3)
	// We can also omit datatype here
	var (
		foo4 = "foo"
		bar4 = "bar"
		baz4 = "baz"
	)
	fmt.Printf("Multiple declaration of foo, bar and baz inside a block without specifying datatype and their corresponding values are: %s, %s and %s\n", foo4, bar4, baz4)

	// 4. Shorthand
	// we can omit datatype as well as 'var' keyword
	foo5 := "foo"
	fmt.Printf("Using shorthand for %s\n", foo5)

	// 5. Constant
	// const <name> <datatype> = <value>
	const c int = 4
	fmt.Printf("c is a constant and its value is %v\n", c)

	// Numeric types
	// Go has several built-in integer types of varying sizes for storing signed and unsigned integers
	// Go has two additional integer types 'byte' and 'rune' that are aliases for 'uint8' and 'int32'
	// In Go there are 2 complex types.
	// 'Complex128' where both real and imaginary parts are float64
	// 'Complex64' where both real and imaginary parts are float32
	c1 := complex(10, 5)
	c2 := 10 + 5i
	fmt.Printf("%v\n", c1)
	fmt.Printf("%v\n", c2)

	// Type conversion
	f := 10.4
	fmt.Printf("Type of f is '%T' and its value is %v\n", f, f)
	i := int(f)
	fmt.Printf("Type of i is '%T' and its value is %v\n", i, i)

}
