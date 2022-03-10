package main

import "fmt"

// Liskov Substitution Principle

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
	GetArea() int
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
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

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

func (r *Square) GetArea() int {
	return r.height * r.width
}

func UseIt(sized Sized) {
	expectedArea := sized.GetArea()
	fmt.Println(fmt.Sprintf("%v area: %d", sized, expectedArea))
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)
}
