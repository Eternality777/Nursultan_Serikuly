package main

import (
	"fmt"
	"github.com/Eternality777/Nursultan_Serikuly/Assignment1/Shapes"
)

func main() {
	rect := Shapes.Rectangle{Length: 5, Width: 3}
	circ := Shapes.Circle{Radius: 7}
	sq := Shapes.Square{Length: 4}
	tri := Shapes.Triangle{SideA: 3, SideB: 4, SideC: 5}

	shapesSlice := []Shapes.Shape{rect, circ, sq, tri}

	for _, shape := range shapesSlice {
		fmt.Println("Shape details:")
		Shapes.PrintShapeDetails(shape)
		fmt.Println()
	}
}
