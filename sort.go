package main

import (
	"fmt"
	"sort"
)

func main() {

	name := []string{"Jaminur", "Belal", "Tarun"}
	sort.Strings(name)
	fmt.Println("Sorted name string is :", name)
	/*
		sorting name in descending order
	*/
	sort.Slice(name, func(i, j int) bool {
		return name[j] < name[i]
	})
	fmt.Println(name)
	nums := []int{7, 2, 4}
	sort.Ints(nums)
	fmt.Println("Ints:   ", nums)

	s := sort.IntsAreSorted(nums)
	fmt.Println("Sorted: ", s)
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] < nums[i]
	})
	fmt.Println("Nums array after sorting in dec order : ", nums)
}
