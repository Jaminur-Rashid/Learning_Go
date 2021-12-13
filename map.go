// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	student := make(map[int]string)
	student[1] = "Jamiunur"
	student[2] = "rashid"
	fmt.Println(student[1])
	for k, v := range student {
		fmt.Println("Key is : ", k, " Value is ", v)
	}
	mpOne := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	//mpOne["Jaminur"] = 1
	for key, value := range mpOne {
		fmt.Println("Key is : ", key, " Value is ", value)
	}
	// check the existence of an key in the map
	_, checkKey := mpOne["a"]
	fmt.Println(checkKey)
}
