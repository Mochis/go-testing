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
}