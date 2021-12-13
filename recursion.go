package main

import (
	"fmt"
)

func print(num int) {
	if num == 0 { //base case
		return
	} else {
		num--
		fmt.Println("Jaminur Rashid")
		print(num)
	}
}
func fibo(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fibo(n-1)
	}
}
func main() {
	print(10)
	fmt.Println(fibo(5))
}
