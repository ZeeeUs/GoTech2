package main

type Pepperoni struct {
	Dough   string
	Sauce   string
	Topping string
}

func newPeperoni() *Pepperoni {
	return &Pepperoni{}
}

func (p *Pepperoni) setDough() {
	p.Dough = "Тонкое тесто"
}

func (p *Pepperoni) setSauce() {
	p.Sauce = "Острый"
}

func (p *Pepperoni) setTopping() {
	p.Topping = "Моцарелла, пеперони, перец чили"
}

func (p *Pepperoni) getPizza() Pizza {
	return Pizza{
		dough: p.Dough,
		sauce: p.Sauce,
		topping: p.Topping,
	}
}
