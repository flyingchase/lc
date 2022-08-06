package main

import (
	"findfeeling"
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("findfeeling.Search(nums, 5): %v\n", findfeeling.Search(nums, 5))
	fmt.Printf("sort.SearchInts(nums, 5): %v\n", sort.SearchInts(nums, 5))
	// fmt.Printf("findfeeling.Quicksort(nums): %v\n", findfeeling.Quicksort(nums))
	fmt.Printf("findfeeling.Heapsort(nums): %v\n", findfeeling.Heapsort(nums))
	fmt.Printf("findfeeling.MergeSort(nums): %v\n", findfeeling.MergeSort(nums))
}
