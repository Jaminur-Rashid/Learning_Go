/*
    Author : Jaminur Rasid
	Date   : 13-12-2021
*/
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i == 20 {
			break
		} else if i%2 == 0 {
			fmt.Println(i, "is a even number")
		} else {
			fmt.Println(i, "is a odd number")
		}
	}
	/*
	While Loop
	*/
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Sum of first thousand numbers is : ",sum)
}
