/*
    Author : Jaminur Rasid
	Date   : 14-12-2021
*/
package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	height float64
	width  float64
}
type Circle struct {
	radius float64
}

/*
define an interface
*/
type geometry interface {
	getArea() float64
	getPerim() float64
}

/*
function that return area of a rectangle
*/
func (r Rectangle) getArea() float64 {

	return r.height * r.width
}

/*
function that return perim of a rectangle
*/
func (r Rectangle) getPerim() float64 {
	return 2*r.height + 2*r.width
}

/*
function that return area of a circle
*/
func (c Circle) getArea() float64 {
	fmt.Println("Function called received")
	return math.Pi * c.radius * c.radius
}

/*
function that return perim of a circle
*/
func (c Circle) getPerim() float64 {
	return 2 * math.Pi * c.radius
}
/*
   we were using seperate methods for Rectangle and circle
   Instead of this we can have used an interface and define
   all the methods for circle and rectangle inside the interface
*/
func (cal Circle) showArea() float64 {// without interface
	fmt.Println("Function called")
	area := cal.getArea()
	fmt.Println(area)
	return area
}

func (cal Rectangle) showArea() float64 {// without interface
	fmt.Println("Function called For Rectangle")
	area := cal.getArea()
	fmt.Println(area)
	return area
}

/*
method to show circle and rectangle details using interface
*/
func showDetails(g geometry) {
	fmt.Println("value is : ", g)
	fmt.Println("Area Ofis : ", g.getArea())
	fmt.Println("Perim is : ", g.getPerim())
}
func main() {
	fmt.Println("Interface Example")
	circleOne := Circle{radius: 3}
	rectangleOne := Rectangle{height: 3, width: 6}
	showDetails(circleOne)
	showDetails(rectangleOne)
	/*
		find the area of Circle
	*/
	fmt.Println(circleOne.showArea())
	/*
		find the area of Rectangle
	*/
	fmt.Println(rectangleOne.showArea())

}
