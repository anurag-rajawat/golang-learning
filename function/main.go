package main

import (
	"fmt"
)

func main() {
	res, err := divide(10, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result =", res)

	//res2, err2 := divide(10, 0)
	//if err2 != nil {
	//	fmt.Println(err2)
	//	return
	//}
	//fmt.Println("Result =", res2)

	//fmt.Println(fnNamedReturn("Anurag"))
	//fmt.Println(add(2, 3, 6))

	// defer
	// 'defer' keyword lets us postpones the execution of a function until the surrounding function returns
	// so, defer is incredibly useful and is commonly used for doing cleanup or error handling.
	//defer fmt.Println("I am finished")
	//fmt.Println("Doing some work...")

}

// Simple declaration
func fn() {}

// With parameter
func fnWithParam(num int) {}

// Shorthand declaration if the consecutive parameters have the same type
func fnShortHand(num1, num2 int) {}

// Go has something special to 'return'
// and that is we can return 2 values from a fn, isn't it cool?
// func <name_of_fn>(<params>) <return_type> {...}
// if want to return multiple values
// func <name_of_fn>(<params>) (<return_type>, <return_type>) {..}
func isOdd(num int) bool {
	return (num & 1) == 0
}

func divide(num1, num2 int) (int, error) {
	if num2 == 0 {
		//return 0, errors.New("division by zero")
		return 0, fmt.Errorf("cannot divide %d by zero", num1)
	}
	return num1 / num2, nil
}

// Named return
// Another cool feature is named returns, where return values can be named and treated as their own variables
// use it with care as this might reduce readability for larger functions
func fnNamedReturn(name string) (s string) {
	s = fmt.Sprintf("Hello %s", name)
	return
}

// Function as values
// In Go functions are first class citizens, and we can use them as values
func fnAsVal() {
	fn := func() {
		fmt.Println("I'm inside fn")
	}
	fn()
}

// Anonymous function
func anonymousFn() {
	func() {
		fmt.Println("I'm an anonymous function")
	}()
}

// Variadic function
func add(val ...int) int {
	sum := 0
	for _, val := range val {
		sum += val
	}
	return sum
}
