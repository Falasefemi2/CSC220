package main

import "fmt"

func main() {
	// Student names
	studentNames := []string{
		"ayo", "bola", "rufai", "jeebolar", "taofeeq",
		"muhammed", "boye", "ajayi", "maryam", "hassan",
	}

	// Student GPAs
	studentGp := []float64{
		3.56, 4.56, 4.32, 5.00, 2.45,
		3.56, 2.45, 2.56, 1.34, 4.34,
	}

	//size of the arrays
	fmt.Println("Number of students:", getArraySize(studentNames))
	fmt.Println("Number of GPAs:", getArraySize(studentGp))
	fmt.Println()

	//elements of the arrays
	fmt.Println("List of student names:")
	getArrayList(studentNames)
	fmt.Println("\nList of student GPAs:")
	getArrayList(studentGp)
	fmt.Println()

	// Search for a student
	searchElementPrint(studentNames, "jeebolar")
	searchElementPrint(studentNames, "femi")
	fmt.Println()

	// Delete a student
	studentNames = deleteElement(studentNames, "maryam")
	fmt.Println("Updated list of student names after deletion:")
	getArrayList(studentNames)
}

// Get the size of an array/slice
func getArraySize[T any](array []T) int {
	return len(array)
}

// Print all elements of an array/slice
func getArrayList[T any](array []T) {
	for _, v := range array {
		fmt.Println(v)
	}
}

// Search for an element and print result
func searchElementPrint[T comparable](array []T, item T) {
	for i, v := range array {
		if v == item {
			fmt.Printf("%v found at index %d\n", item, i)
			return
		}
	}
	fmt.Printf("%v not found in the array\n", item)
}

// Delete an element from a slice
func deleteElement[T comparable](array []T, item T) []T {
	for i, v := range array {
		if v == item {
			return append(array[:i], array[i+1:]...)
		}
	}
	fmt.Printf("%v not found, nothing deleted\n", item)
	return array
}
