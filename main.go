package main

import (
	"fmt"
)

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	arr := []int{}
	if len(slice1) == len(slice2) {
		for i := len(slice2) - 1; i >= 0; i-- {
			arr = append(arr, slice1[i])
			arr = append(arr, slice2[i])

		}
	}
	if len(slice1) > len(slice2) {
		for i := len(slice1) - 1; i >= 0; i-- {
			if i > len(slice2)-1 {
				arr = append(arr, slice1[i])
			} else {
				arr = append(arr, slice1[i])
				arr = append(arr, slice2[i])
			}
		}
	}
	if len(slice2) > len(slice1) {
		for i := len(slice2) - 1; i >= 0; i-- {
			if i > len(slice1)-1 {
				arr = append(arr, slice2[i]) 
			} else {
				arr = append(arr, slice1[i])
				arr = append(arr, slice2[i])
			}
		}
	}

	return arr
}
