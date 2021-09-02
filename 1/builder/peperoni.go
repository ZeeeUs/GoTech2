package main

type Peperoni struct {
	Dough   string
	Sauce   string
	Topping string
}

func newPeperoni() *Peperoni {
	return &Peperoni{}
}

func (p *Peperoni) makeDough(dough string) {
	p.Dough := dough
}

func (p *Peperoni) makeSauce(sauce string) {
	p.Sauce := sauce
}

func (p *Peperoni) makeTopping(topping string) {
	p.Topping := topping
}

func (p *Peperoni) Build() Pizza {
	return Pizza{
		dough: p.Dough,
		sauce: p.Sauce,
		topping: p.Topping,
	}
}
