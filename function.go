/*
    Author : Jaminur Rasid
	Date   : 13-12-2021
*/
package main

import "fmt"

func getUserDetails(name string, id string) string {
	return (name + " " + id)
}

/*
    function that returns sum of all elements
	of an array
*/
func calTotal(arr [6]int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

/*
   multiple return values in function
*/
func returnMultiple() (int, string) {
	return 34, "jaminur Rashid"
}

/*
variadic function
-can be called with any number of trailing arguments
*/
func calSum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
func main() {
	userOne := getUserDetails("Jaminur Rashid", "CE-16034")
	fmt.Println("User One details is : ", userOne)
	// passing array to a function an a parameter
	arr := [6]int{1, 2, 3, 5, 6, 7}
	totSum := calTotal(arr)
	fmt.Println(totSum)
	valOne, valTwo := returnMultiple()
	fmt.Println("Vlue one is : ", valOne, " Value two is : ", valTwo)
	/*
		variadic function
	*/
	sum := calSum(1, 2, 4)
	fmt.Println("Sum is :", sum)
}
