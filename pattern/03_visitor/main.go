package main

import "fmt"

type Bank interface {
	getType() string
	accept(visitor)
}

type visitor interface {
	visitSber(*Sber)
	visitAlfa(*Alfa)
}

type Sber struct {
	Dollar float64
}

func (s *Sber) getType() string {
	return "Sberbank"
}

func (s *Sber) getCurrency() float64 {
	return 75.9
}

func (s *Sber) accept(v visitor) {
	v.visitSber(s)
}

type Alfa struct {
	Dollar float64
}

func (a *Alfa) getType() string {
	return "AlfaBank"
}

func (a *Alfa) getCurrency() float64 {
	return 86.9
}

func (a *Alfa) accept(v visitor) {
	v.visitAlfa(a)
}

type calculateDollarToRub struct {
	count float64
}

func (dtr *calculateDollarToRub) visitSber(s *Sber) {
	fmt.Println(dtr.count * s.getCurrency())
}

func (dtr *calculateDollarToRub) visitAlfa(a *Alfa) {
	fmt.Println(dtr.count * a.getCurrency())
}

func main() {
	Sber := &Sber{Dollar: 75.8}
	Alfa := &Alfa{Dollar: 76.0}

	calculateDollarToRub := &calculateDollarToRub{count: 3}

	fmt.Println(Sber.getType())
	Sber.accept(calculateDollarToRub)
	fmt.Println(Alfa.getType())
	Alfa.accept(calculateDollarToRub)

}
