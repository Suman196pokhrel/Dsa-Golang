package arraysgo

import "fmt"

func ArrayOperations() {
	// integer Array
	arr := []int{}
	element := 11

	// INSERTING ELEMENTS
	addElements(&arr, element)
	addElements(&arr, element+1)
	addElements(&arr, element+2)
	addElements(&arr, element+3)
	fmt.Println("Inserting Elements in array => ", arr)

	// UPDATING ELEMENTS
	updateElements(&arr, 0, 55)
	updateElements(&arr, 1, 65)
	updateElements(&arr, 2, 75)
	updateElements(&arr, 3, 85)
	fmt.Println("Updating Elements in array => ", arr)

	// DELETE ELEMENTS
	deleteElements(&arr, 0)
	fmt.Println("Deleting Elements in array => ", arr)

}

func addElements(arr *[]int, element int) {
	*arr = append(*arr, element)
}

func updateElements(arr *[]int, index int, element int) {
	(*arr)[index] = element
}

func deleteElements(arr *[]int, index int) {
	*arr = append((*arr)[:index], (*arr)[index+1:]...)
}
