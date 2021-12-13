package main

import "fmt"

// Creating structure
type Student struct {
	name   string
	branch string
	year   int
}
type Thesis struct {
	name   string
	author string
	year   int
}

// Creating nested structure
type Teacher struct {
	name    string
	subject string
	exp     int
	thesis  Thesis  // Thesis structure
	details Student // Student structure
}

func main() {

	// Initializing the fields
	// of the structure
	result := Teacher{
		name:    "Suman",
		subject: "Java",
		exp:     5,
		thesis:  Thesis{"Line following car", "jaminur rashid", 2010},
		details: Student{"Bongo", "CSE", 2},
	}
	fmt.Println(result.details)
	fmt.Println(result)
}
