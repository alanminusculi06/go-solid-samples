package main

import "fmt"

// Interface Segregation Principle
// Uma classe não deve ser forçada a implementar interfaces e métodos que não irão utilizar

type Animal interface {
	Walk()
	Swim()
}

type Dog struct{ Name string }

func (dog Dog) Walk() {
	fmt.Println(dog.Name + " is walking")
}

func (dog Dog) Swim() {
	fmt.Println(dog.Name + " is swimming")
}

func main() {
	dog := Dog{Name: "Beethoven"}
	toWalk(dog)
	toSwim(dog)
}

func toWalk(animal Animal) {
	animal.Walk()
}

func toSwim(animal Animal) {
	animal.Swim()
}
