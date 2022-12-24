package main

import "fmt"

func main() {
	// A struct is a user-defined type that contains a collection of named fields.
	// If you're coming from an objected-oriented background,
	// think of structs as lightweight classes which that support composition but not inheritance.
	// Two structs are equal if all of their corresponding fields are equal.
	// if a struct field is declared with a lower case identifier,
	// it will not be exported and only be visible to the package it is defined in.
	// type <name> struct {...}

	// One way of initializing a struct
	h := Human{Name: "Anurag", Age: 19}
	// We use '.' (period) to access struct fields
	fmt.Println("My name is", h.Name)
	fmt.Println("My age is", h.Age)

	// Another way to initializing a struct, in this way a trailing comma ',' is required
	h1 := Human{
		Name: "Anurag",
		Age:  19,
	}
	fmt.Println("My name is", h1.Name)
	fmt.Println("My age is", h1.Age)

	// We can also initialize a subset of fields.
	// In this case the remaining fields get initialized with their zero values.
	h2 := Human{
		Name: "Unknown",
	}
	fmt.Println("My name is", h2.Name)
	fmt.Println("My age is", h2.Age) // This will give 0 (Zero value)

	// An example of embedded struct
	st := Student{}
	st.Id = 1
	st.Name = "Revat"
	st.Age = 20

	fmt.Println("Student's name =>", st.Name)
	fmt.Println("Student's Age =>", st.Age)
	fmt.Println("Student's ID =>", st.Id)

	// Anonymous struct
	a := struct {
		Name string
	}{"Missing-name"}

	fmt.Println()
	fmt.Println(a.Name)
	st1 := Student{
		Human: Human{Name: "Anurag", Age: 19},
		Id:    1,
	}
	fmt.Println(st1.Name)

	// This is an example of composite struct
	h3 := Human{
		Name: "Sanyog",
		Age:  19,
	}
	s1 := St{
		H:  h3,
		Id: 3,
	}

	fmt.Println(s1.Id)
	fmt.Println(s1.H.Name)
	fmt.Println(s1.H.Age)

	// Method example
	car := Car{
		Name:  "Jaguar F-Pace",
		Price: 9000000,
		Year:  2021,
	}
	fmt.Printf("Price of %v is ₹%v\n", car.Name, car.Price)
	//isLatest := car.IsLatest()
	//if isLatest {
	//	fmt.Println(car.Name, "is a latest car.")
	//} else {
	//	fmt.Println(car.Name, "is not a latest car.")
	//}

	car.UpdatePrice(10000000)
	fmt.Printf("New price of %v is ₹%v\n", car.Name, car.Price)
}

type Human struct {
	Name string
	Age  int
}

// Student is an embedded struct
// we can think embedding like inheritance
// But it is not preferred, composition is preferred.
type Student struct {
	Human
	Id int
}

// St is a composite struct
type St struct {
	H  Human
	Id int
}

// Animal struct with tags
// A struct tag is just a tag that allows us to attach metadata information
// to the field which can be used for custom behavior using the reflect package.
type Animal struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Car struct {
	Name  string
	Year  int
	Price int
}

// IsLatest is a Method (function receivers)
// Method is used to define a function on a type
// think it as similar to method in a class in object-oriented language
// we can access the instance of Car using the receiver variable c.
// Think of it as 'this' keyword from the object-oriented world.
func (c *Car) IsLatest() bool {
	return c.Year >= 2015
}

func (c *Car) UpdatePrice(price int) {
	c.Price = price
}
