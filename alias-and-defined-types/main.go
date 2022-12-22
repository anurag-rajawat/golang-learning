package main

import "fmt"

// They allow developers to provide an alternate name for an existing type and use it interchangeably with the underlying type.
// type <alias> = <existing_type>
type str = string

// We have defined types that unlike alias types do not use an equals sign.
type defined int64

// What's the difference?
// So, defined types do more than just give a name to a type.
// It first defines a new named type with an underlying type. However, this defined type is different from any other type, including its underline type.
// Hence, it cannot be used interchangeably with the underlying type like alias types.

func main() {
	var alias str = "I'm an Alias"
	fmt.Printf("'alias' type is %T and its value is %v\n", alias, alias)

	var d defined = 2
	fmt.Printf("'alias' type is %T and its value is %v\n", d, d)
}
