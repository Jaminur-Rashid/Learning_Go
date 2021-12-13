/*
    Author : Jaminur Rasid
	Date   : 13-12-2021
*/
package main

import "fmt"

func main() {
	marks := 70
	if marks >= 60 {
		fmt.Println("Cgpa is : A-")
	} else if marks >= 70 && marks <= 75 {
		fmt.Println("Cgpa is A")
	} else if marks < 40 {
		fmt.Println("Failed")
	} else {
		fmt.Println("Cgpa is :  A+ ")
	}
}
