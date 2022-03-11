package main

import "fmt"

// Liskov Substitution Principle

type Sized interface {
	GetArea() int
}

type Rectangle struct {
	width, height int
}

func NewRectangle(width, height int) Sized {
	return &Rectangle{width, height}
}

func (r *Rectangle) GetArea() int {
	return r.height * r.width
}

// Se uma funcao recebe um interface e
// funciona com um tipo T que implementa a interface
// qualquer estrutra compatível com T
// também deve ser capaz de usar a funcao
type Square struct {
	Rectangle
}

func NewSquare(size int) Sized {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (r *Square) GetArea() int {
	return r.height * r.width
}

func UseIt(sized Sized) {
	area := sized.GetArea()
	fmt.Println(fmt.Sprintf("%v area: %d", sized, area))
}

func main() {
	rc := NewRectangle(2, 3)
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)
}
