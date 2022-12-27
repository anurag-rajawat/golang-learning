package main

import "fmt"

func main() {

	// Unlike other languages, Go Interfaces are implemented implicitly, so we don't need something like an implements' keyword.
	// This means, that a type satisfies an interface automatically when it has "all the methods" of the interface.
	// We can embed interface like structs
	// Interface values are comparable

	m1 := Mobile{
		Brand: "Apple",
		Power: 10,
	}

	m2 := Mobile{
		Brand: "Redmi",
		Power: 33,
	}

	s := Socket{}

	s.PlugIn(m1, m1.Power)
	s.PlugIn(m2, m2.Power)

	l := Laptop{Brand: "HP", CPU: "Ryzen 3", Power: 65}
	l2 := Laptop{
		Brand: "Apple",
		CPU:   "M1 Pro",
		Power: 96,
	}

	s.PlugIn(l, l.Power)
	s.PlugIn(l2, l2.Power)

	// Empty interface
	// Empty interfaces can be used to handle values of unknown type.
	// Some examples are:
	// Reading heterogeneous data from an API
	// Variables of an unknown type, like in the fmt.Prinln function
	// To use a value of type empty interface{}, we can use type assertion or a type switch to determine the type of the value.
	var v interface{}
	fmt.Println(v)

	// A type assertion provides access to an interface value's underlying concrete value.
	// Type switch
	// A switch statement can be used to determine the type of variable of type empty interface{}.
	var i interface{} = "Hello!"
	switch i.(type) {
	case int:
		fmt.Printf("int: %d\n", i)
	case string:
		fmt.Printf("string: %s\n", i)
	case bool:
		fmt.Printf("bool: %v\n", i)
	default:
		fmt.Printf("Unknown\n")
	}
}

type PowerDrawer interface {
	Draw(power int)
}

type Mobile struct {
	Brand string
	Power int
}

type Laptop struct {
	Brand string
	CPU   string
	Power int
}

type Socket struct {
}

func (s Socket) PlugIn(device PowerDrawer, power int) {
	device.Draw(power)
}

func (m Mobile) Draw(power int) {
	fmt.Printf("Mobile Phone of brand \"%v\" uses %v units of power.\n", m.Brand, power)
}

func (l Laptop) Draw(power int) {
	fmt.Printf("Laptop of brand \"%v\" having \"%v\" CPU uses %v units of power.\n", l.Brand, l.CPU, power)
}
