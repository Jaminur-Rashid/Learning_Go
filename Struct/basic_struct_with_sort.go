package main

import (
	"fmt"
	"sort"
)

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
	personThree := person{name: "Tarun", country: "BD", age: 40}
	personTwo := person{"Jaminur Rashid", "Bangladesh", 25}
	fmt.Println("personTwo of type person ", personTwo)
	fmt.Println(personTwo.name)
	fmt.Println(personThree)
	/*
		Pointers to a struct
	*/
	personFour := &person{"Belal Hossain", "BD", 12}
	fmt.Println(&personFour)
	fmt.Println((*personFour).name)
	personArr := []person{personThree, personTwo, personOne, *personFour}
	fmt.Println("Before Sorting")
	fmt.Println(personArr)
	/*
		sort structure by properties name here we have sorted the structure
		person by age in ascending order. before that we have pushed all the
		person into an slice then sorted it by age
	*/
	sort.Slice(personArr, func(i, j int) bool {
		return personArr[i].age < personArr[j].age
	})
	fmt.Println("After Sorting")
	fmt.Println(personArr[0])
	fmt.Println(personArr)

}
