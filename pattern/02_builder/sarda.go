package main

type Sarda struct {
	dough string
	sauce string
	topping string
}

func newSarda() *Sarda{
	return &Sarda{}
}

func (s *Sarda) setDough() {
	s.dough = "Классическое тесто"
}

func (s *Sarda) setSauce() {
	s.sauce = "Не острый"
}

func (s *Sarda) setTopping() {
	s.topping = "Томатный соус, моцарелла, сыр пекорино и острая салями"
}

func (s *Sarda) getPizza() Pizza {
	return Pizza{
		dough: s.dough,
		sauce: s.sauce,
		topping: s.topping,
	}
}