package main

import "errors"

// Liskov Substitution Principle
// Uma classe derivada deve ser substituível por sua classe base.

// Se uma funcao recebe uma interface e
// funciona com um tipo T que implementa a interface
// qualquer estrutra compatível com T
// também deve ser capaz de usar a funcao
type Vehicle interface {
	StartUp() error
}

type Car struct {
	Model     string
	FuelLevel int
}

func (c Car) StartUp() error {
	if c.FuelLevel < 5 {
		return errors.New("fuel level too low")
	}
	return nil
}

type HybridCar struct {
	Car
	BatteryLevel int
}

func (h HybridCar) StartUp() error {
	if h.FuelLevel < 5 && h.BatteryLevel < 5 {
		return errors.New("fuel and battery levels too low")
	}
	return nil
}

func main() {
	car := Car{Model: "ABC GT", FuelLevel: 4}
	err := StartVehicle(car)
	if err != nil {
		panic(err)
	}
}

func StartVehicle(v Vehicle) error {
	return v.StartUp()
}
