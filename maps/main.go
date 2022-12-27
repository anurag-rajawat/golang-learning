package main

import "fmt"

func main() {
	// Maps are reference types, which means when we assign a map to a new variable they both refer to the same underlying data structure.
	// Therefore, changes done by one variable will be visible to the other.

	// Declaration and Initialization

	// using make function
	var marks map[string]int = make(map[string]int)
	var marks1 = make(map[string]int)

	// map literal, last trailing comma is necessary
	var m = map[string]int{
		"anurag": 1,
		"revat":  2,
	}
	fmt.Println(m)

	// using shorthand notation
	marks2 := make(map[string]int)

	fmt.Println(marks)
	fmt.Println(marks1)
	fmt.Println(marks2)

	// Retrieving value
	// When you retrieve the value assigned to a given key, it returns an additional boolean value as well.
	// The boolean variable will be true if the key exists, and false otherwise.
	val, ok := marks2["anurag"]
	if ok {
		fmt.Println(val)
	}

	// Updating value of a key
	marks2["anurag"] = 99
	//marks2["anurag"] = 93
	marks2["revat"] = 100 // Maths
	marks2["astha"] = 99
	mark, ok := marks2["anurag"]
	if ok {
		fmt.Println(mark)
	}

	// Deleting a key
	delete(marks2, "anurag")
	mark, ok = marks2["anurag"]
	if ok {
		fmt.Println(mark)
	}

	// Iteration
	for student, marks := range marks2 {
		fmt.Printf("%v's marks are %v\n", student, marks)
		fmt.Println(student+"'s", "marks are", marks)
	}

	dummyMap := marks2
	dummyMap["missing"] = 100
	delete(dummyMap, "revat")
	fmt.Println("Dummy Map => ", dummyMap)
	fmt.Println("Original Map => ", marks2)

}
