package Shapes

import "math"

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Length float64
    Width  float64
}

type Circle struct {
    Radius float64
}

type Square struct {
    Length float64
}

type Triangle struct {
    SideA, SideB, SideC float64
}

func (r Rectangle) Area() float64 {
    return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Length + r.Width)
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

func (s Square) Area() float64 {
    return s.Length * s.Length
}

func (s Square) Perimeter() float64 {
    return 4 * s.Length
}

func (t Triangle) Area() float64 {
    
    s := (t.SideA + t.SideB + t.SideC) / 2 
    area := math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
    return area
}

func (t Triangle) Perimeter() float64 {
    return t.SideA + t.SideB + t.SideC
}
