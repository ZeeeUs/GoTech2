package main

import "fmt"

func main() {
	pepperoniBuilder := getBuilder("Pepperoni")
	sardaBuilder := getBuilder("Sarda")

	director := newDirector(pepperoniBuilder)
	pepperoniPizza := director.makePizza()

	fmt.Printf("Dough type: %s\n", pepperoniPizza.dough)
	fmt.Printf("Sauce type: %s\n", pepperoniPizza.sauce)
	fmt.Printf("Topping: %s\n", pepperoniPizza.topping)

	director.setBuilder(sardaBuilder)
	sardaPizza := director.makePizza()

	fmt.Printf("Dough type: %s\n", sardaPizza.dough)
	fmt.Printf("Sauce type: %s\n", sardaPizza.sauce)
	fmt.Printf("Topping: %s\n", sardaPizza.topping)
}
