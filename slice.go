/*
    Author : Jaminur Rasid
	Date   : 13-12-2021
*/
package main

import "fmt"

func main() {
	arr := make([]int, 100)
	for i := 0; i < 10; i++ {
		//fmt.Println(arr[i])
		arr[i] = i + 10
	}
	// Before append
	for i := 0; i < 15; i++ {
		fmt.Println("Value at index ", i, "is ", arr[i])
	}
	arr = append(arr, 11)
	fmt.Println("After append")
	//after append
	for i := 0; i < 15; i++ {
		fmt.Println("Value at index ", i, "is ", arr[i])
	}
	arrOne := arr[2:10] // copy elements of arr to arrOne from index 2 to 9
	for i := 0; i < len(arrOne); i++ {
		fmt.Println(" value at index ", i, "is ", arrOne[i])
	}
	arrTwo := arrOne[2:] // copy all elements from index two
	for i := 0; i < len(arrTwo); i++ {
		fmt.Println(" value at index ", i, "is ", arrTwo[i])
	}
	arrOfString := make([]string, 10)
	arrOfString[0] = "a"
	arrOfString[1] = "b"
	arrOfString[2] = "c"
	for i := 0; i < len(arrOfString); i++ {
		fmt.Println(" value at index ", i, "is ", arrOfString[i])
	}
}
