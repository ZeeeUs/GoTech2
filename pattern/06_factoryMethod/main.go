package main

import "fmt"

type Engine interface {
	getAutoName() string
	setType(engType string)
	getType() string
	setPower(pow int)
	getPower() int
}

type Auto struct {
	name     string
	engType  string
	engPower int
}

func (a *Auto) setType(engType string) {
	a.engType = engType
}

func (a *Auto) getType() string {
	return a.engType
}

func (a *Auto) setPower(engPower int) {
	a.engPower = engPower
}

func (a *Auto) getPower() int {
	return a.engPower
}

func (a *Auto) getAutoName() string {
	return a.name
}

type Ford struct {
	Auto
}

func newFord() Engine {
	return &Ford{
		Auto{
			name:     "Ford Transit",
			engType:  "Diesel",
			engPower: 234,
		},
	}
}

type Mercedes struct {
	Auto
}

func newMercedes() Engine {
	return &Mercedes{
		Auto{
			name:     "Mercedes E-class",
			engType:  "Gas",
			engPower: 315,
		},
	}
}

func getAuto(auto string) (Engine, error) {
	if auto == "Ford" {
		return newFord(), nil
	}
	if auto == "Mercedes" {
		return newMercedes(), nil
	}

	return nil, fmt.Errorf("Wrong auto type")
}

func printInfo(e Engine) {
	fmt.Printf("Name: %s\n", e.getAutoName())
	fmt.Printf("Engine type: %s\n", e.getType())
	fmt.Printf("Engine power: %v\n", e.getPower())
}

func main() {
	ford, _ := getAuto("Ford")
	mercedes, _ := getAuto("Mercedes")

	printInfo(ford)
	printInfo(mercedes)
}
