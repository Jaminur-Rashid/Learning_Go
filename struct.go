package main

import "fmt"

type person struct {
	name    string
	country string
	age     int
}

func main() {
	fmt.Println("Struct Example ")
	/*
		Declaring a variable of a `struct` type
		All the struct fields are initialized
		with their zero value
	*/
	var personOne person
	fmt.Println(personOne)
	personTwo := person{"Jaminur Rashid", "Bangladesh", 25}
	fmt.Println("personTwo of type person ", personTwo)
	fmt.Println(personTwo.name)
	personThree := person{name: "Tarun", country: "BD"}
	fmt.Println(personThree)
	/*
		Pointers to a struct
	*/
	personFour := &person{"Belal Hossain", "BD", 25}
	fmt.Println(&personFour)
	fmt.Println((*personFour).name)
}
