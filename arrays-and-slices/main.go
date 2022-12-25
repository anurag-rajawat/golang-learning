package main

import "fmt"

func main() {
	// Arrays
	// Arrays in Go are value types unlike other languages like C, C++, and Java where arrays are reference types.
	// This means that when we assign an array to a new variable or pass an array to a function, the entire array is copied.
	var nums = [5]int{1, 2, 3, 4, 5}
	var nums2 = []int{1, 2, 3, 4, 5}
	nums3 := []int{1, 2, 3}

	_ = len(nums2)
	_ = len(nums3)

	// Traditional way of iterating
	for i := 0; i < len(nums); i++ {
		fmt.Printf("Element at index %v = %v\n", i, nums[i])
	}

	// Another way of iterating
	// for _, e := range arr {} // Omit index with _ and use element
	// for i := range arr {} // Use index only
	// for range arr {} // Simply loop over the array
	for i, v := range nums {
		fmt.Printf("Element at index %v = %v\n", i, v)
	}

	// Multi-dimensional
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("Element at index [%v][%v] is %v\n", i, j, matrix[i][j])
			fmt.Printf("%v ", matrix[i][j])
		}
		fmt.Println()
	}

	for i, v := range matrix {
		fmt.Printf("Array at index %v is %v\n", i, v)
	}

	for i, v := range matrix {
		for j, val := range v {
			fmt.Printf("Element at index [%v][%v] is %v\n", i, j, val)
		}
	}
	nums4 := []int{1, 2, 3, 4, 5}
	Reverse(nums4)
	fmt.Println(nums4)

	// Slice
	// We can think slice as dynamic arrays like ArrayList for Java People
	// Slices are reference types, unlike arrays.
	// One way of initializing
	a := []int{20, 15, 5, 30, 25}
	b := a[1:4]
	fmt.Println(b)
	s := make([]int, 0, 10)
	fmt.Println(s == nil)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	langs := []string{"Java", "Go"}
	toLearnLangs := make([]string, len(langs))
	_ = copy(toLearnLangs, langs)
	fmt.Println(toLearnLangs)
	langs = append(langs, "JavaScript", "TypeScript", "Python")
	fmt.Println(langs)

	// Another way using 'make' function
	remainingSubj := make([]string, 0)
	remainingSubj = append(remainingSubj, "DB", "COA")
	fmt.Println(remainingSubj)
}

func Reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
		i++
		j--
	}
}
