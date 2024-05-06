package sorting

import "fmt"

/*
function to demonstrate a standard bubble sort algorithm that takes an
address of an array , with elements of integer data type and returns
the address sorted array
*/
func BubbleSort() {
	var arr = []int{22, 33, 1, 23, 5, 3, 4}

	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// SWAP VARIABLES
				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
		}
	}

	fmt.Println(arr)

}
