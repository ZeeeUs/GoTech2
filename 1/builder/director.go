package main

type director struct {
	builder PizzaBuilder
}

func newDirector(pb PizzaBuilder) *director {
	return &director{
		builder: pb,
	}
}

func (d *director) setBuilder(pb PizzaBuilder) {
	d.builder = pb
}

func (d *director) makePizza() Pizza {
	d.builder.setDough()
	d.builder.setSauce()
	d.builder.setTopping()
	return d.builder.getPizza()
}