package main

import "fmt"

func main() {
	// 1. if/else
	// Traditional way
	marks := 99
	if marks < 33 {
		fmt.Println("Sorry to say you've been failed!")
	} else if marks < 80 {
		fmt.Println("You get C grade")
	} else if marks < 90 {
		fmt.Println("You get B grade")
	} else if marks > 90 {
		fmt.Println("You get A grade")
	} else {
		fmt.Println("Aap is gole k nahi ho")
	}

	// Compact if
	if m := 101; m < 33 {
		fmt.Println("Sorry to say you've been failed!")
	} else if m < 80 {
		fmt.Println("You get C grade")
	} else if m < 90 {
		fmt.Println("You get B grade")
	} else if m > 90 && m <= 100 {
		fmt.Println("You get A grade")
	} else {
		fmt.Println("Kaha se aaye ho bhai lagta h aap is gole k nahi ho 😂")
	}

	// 2. Switch
	// Traditional
	// We don't need 'break' statement it's automatically added at the end of each case unlike other languages.
	favSubj := "Math"
	//favSubj := "Chemistry" // most hated subject
	switch favSubj {
	case "English":
		fmt.Println("Common subject")
	case "CS":
		fmt.Println("Second favorite subject")
	case "Math":
		fmt.Println("Favorite subject")
	default:
		fmt.Println("Mujhe nahi pata")
	}

	// Switch using shorthand declaration
	switch fSubj := "CS"; fSubj {
	case "English":
		fmt.Println("Common subject")
	case "CS":
		fmt.Println("Second favorite subject")
	case "Math":
		fmt.Println("Favorite subject")
	default:
		fmt.Println("Mujhe nahi pata")
	}

	// We can also use the fallthrough keyword to transfer control to the next case
	// even though the current case might have matched.
	// I don't know why do we need it?
	//switch day := "monday"; day {
	//case "monday":
	//	fmt.Println("time to work!")
	//	fallthrough
	//case "friday":
	//	fmt.Println("let's party")
	//default:
	//	fmt.Println("browse memes")
	//}

	// 3. Loops
	// In Go there is only one loop which is 'for' nothing like 'while' or 'do while'
	// and 'break' and 'continue' are same as other languages
	for i := 0; i < 10; i++ {
		fmt.Println("Learn Go!")
	}

	// while loop
	for {
		fmt.Println("Learn Go!")
	}
}
