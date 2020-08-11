package main

import "fmt"

func main() {
	arr := [...]int {1, 2, 3, 4, 5}
	// iterate
	for i, v := range(arr) {
		fmt.Printf("Index[%d], Value: %d\n", i, v)
	}
	// slices
	sli := make([]int, 5, 10)
	for i, v := range(sli) {
		fmt.Printf("Slice Index[%d], Value: %d\n", i, v)
	}
	fmt.Println()
	// append to slice
	sli = append(sli, 10)
	for i, v := range(sli) {
		fmt.Printf("Slice Index[%d], Value: %d\n", i, v)
	}
	// changing slice and underlying array
	arr2 := [...]int{1, 2, 3}
	sli2 := arr2[0:2]
	fmt.Printf("arr2[0] before change slice: %d\n", arr2[0])
	sli2[0] = 11
	fmt.Printf("arr2[0] after change slice: %d\n", arr2[0])
}