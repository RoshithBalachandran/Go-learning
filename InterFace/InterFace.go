package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Parameter() float64
}

type Circle struct {
	Radius float64
}

func (C Circle) Area() float64 {
	return math.Pi*C.Radius*C.Radius
}

func (C Circle) Parameter()float64{
	return 2*math.Pi*C.Radius
}

type Rectangle struct {
	Height float64
	width  float64
}

func (R Rectangle) Area()float64{
	return R.width*R.Height
}

func (R Rectangle)Parameter()float64{
	return 2*(R.width+R.Height)
}

func PrintShape(S Shape){
	fmt.Printf("Area is %.2f \n",S.Area())
	fmt.Printf("Parameter is %.2f\n",S.Parameter())
}
func main() {
	circle := Circle{8}
	rectangle := Rectangle{5, 9}
	fmt.Println("Circle :-")
	PrintShape(circle)
	fmt.Println("Rectangle :-")
	PrintShape(rectangle)
}