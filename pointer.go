package main

import "fmt"

func changeNameOne(name string) {
	name = "Tarun"
}
func changeName(name *string) {
	*name = "Tarun"
}

func main() {
	name := "Jaminur Rashid"
	fmt.Println(name)
	changeNameOne(name)
	fmt.Println("Name change without pointer is : ", name)
	changeName(&name)
	fmt.Println("Nafe after changing is : ", name)
	fmt.Println("Address of name is : ", &name)
}
