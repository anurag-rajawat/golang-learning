package main

import (
	"fmt"
)

// In Go there is nothing like exceptions as we've in other languages
// Instead, we can just return a built-in error type which is an interface type.
func main() {
	//ans, err := Divide(4, 0)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(ans)

	// Sentinel errors
	// Another important technique in Go is defining expected Errors, so they can be checked explicitly in other parts of the code.
	// These are sometimes referred to as sentinel errors.
	// In Go, it is considered conventional to prefix the variable with Err.
	//var errDivideByZero = errors.New("cannot divide by zero")

	//ans2, err2 := Divide(4, 1)
	//if err2 != nil {
	//	// we can check explicitly which error occurred using the errors.Is function
	//	switch {
	//	case errors.Is(err2, errDivideByZero):
	//		fmt.Println(err2)
	//	default:
	//		fmt.Println("I have no idea")
	//	}
	//	return
	//}
	//fmt.Println(ans2)

	//ans3, err3 := Divide(10, 0)
	//if err3 != nil {
	//	var divErr DivisionErr
	//	switch {
	//	// 'errors.As' function checks whether the error has a specific type,
	//	// unlike the Is(), which examines if it is a particular error object.
	//	case errors.As(err3, &divErr):
	//		fmt.Println(err3)
	//	default:
	//		fmt.Println("Unknown")
	//	}
	//}
	//fmt.Println(ans3)

	// Panic and Recover

	// There are some situations where the program cannot continue
	// in those cases we can use the built-in 'panic' function
	// it will stop the normal execution of the current goroutine and
	// the control is returned back to the caller.
	// This is repeated until the program exits with the panic message
	// and stack trace.
	// Good thing is that we can regain control on a panicking program
	// using the built-in 'recover' function along with 'defer' keyword
	// we can think them as try/catch idiom in other language,
	// but we should avoid panic and error and use errors when possible

	// We should use 'panic' in two valid cases:
	// 1. An unrecoverable error
	// For example, reading a configuration file which is important to start the program, as there is nothing else to do if the file read itself fails.
	// 2. Developer error
	// For example, de-referencing a pointer when the value is nil will cause a panic
	ans := Divide2(1, 0)
	fmt.Println(ans)
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		// Error construction using 'New' function of errors package
		//return 0, errors.New("cannot divide by zero")

		// Formatted error
		//return 0, fmt.Errorf("cannot divide %v by zero", a)

		// Custom error
		return 0, DivisionErr{
			Code: -1,
			Msg:  "cannot divide by zero",
		}
	}
	return a / b, nil
}

// DivisionErr is a custom error
type DivisionErr struct {
	Code int
	Msg  string
}

func (receiver DivisionErr) Error() string {
	return fmt.Sprintf("code %d: %s", receiver.Code, receiver.Msg)
}

func Divide2(a, b int) int {
	defer HandlePanic()
	if b == 0 {
		msg := fmt.Sprintf("cannot divide %v by zero", a)
		panic(msg)
	}
	return a / b
}

func HandlePanic() {
	data := recover()
	fmt.Printf("Successfully recovered data: %v\n\n", data)
}
