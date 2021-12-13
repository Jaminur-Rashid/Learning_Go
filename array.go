/*
    Author : Jaminur Rasid
	Date   : 13-12-2021
*/
package main

import "fmt"

func main() {
	var numArr [15]int
	numArr[0] = 10
	numArr[1] = 6
	for i := 0; i <= 10; i++ {
		if i == 0 || i == 1 {
			continue
		} else {
			numArr[i] = i
		}
	}
	for i := 0; i <= 14; i++ {
		fmt.Println("Value at index ", i, "is ", numArr[i])
	}
	numArr[4] = 43
	// After mutating value at index 4
	fmt.Println(numArr[4])
	/*
	   Another way of array declaration
	*/
	numArrOne := [10]string{"jaminur", "Belal Hossain"}
	numArrOne[4] = "Tarun"
	fmt.Println(numArrOne[0])
	fmt.Println(numArrOne[4])
	/*
		Two dimentional array
	*/

	twoDimArr := [15][15]int{{0, 1, 2}}

	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			twoDimArr[i][j] = i
		}
	}
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			fmt.Println("Value is : index ", i, " ", j, twoDimArr[i][j])
		}
	}
}
