package main

type PizzaBuilder interface {
	setDough()
	setSauce()
	setTopping()
	getPizza() Pizza
}

func getBuilder(builderType string) PizzaBuilder{
	if builderType == "Pepperoni" {
		return &Pepperoni{}
	}

	if builderType == "Sarda" {
		return &Sarda{}
	}

	return nil
}
