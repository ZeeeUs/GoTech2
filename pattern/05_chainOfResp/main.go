package main

import "fmt"

type checkPoint interface {
	execute(*passenger)
	setNext(checkPoint)
}

type passenger struct {
	name string
	entranceInspectionDone bool
	regOnBoardingDone bool
	customsControlDone bool
}

type entranceInspection struct {
	next checkPoint
}

func (e *entranceInspection) execute(p *passenger) {
	if p.entranceInspectionDone {
		fmt.Println("Пассажир уже прожёл досмотр на входе")
		e.next.execute(p)
		return
	}
	fmt.Println("Досмотр на входе пройден")
	p.entranceInspectionDone = true
	e.next.execute(p)
}

func (e *entranceInspection) setNext(next checkPoint) {
	e.next = next
}

type registration struct {
	next checkPoint
}

func (r *registration) execute (p *passenger) {
	if p.regOnBoardingDone {
		fmt.Println("Пассажир уже прошёл регистрацию на посадку")
		r.next.execute(p)
		return
	}
	p.regOnBoardingDone = true
	fmt.Println("Регистрация пройдена")
	r.next.execute(p)
}

func (r *registration) setNext(next checkPoint) {
	r.next = next
}

type customsControl struct {
	next checkPoint
}


func (c *customsControl) execute (p *passenger) {
	if p.customsControlDone {
		fmt.Println("Пассажир уже прошёл таможенный контроль")
	}
	p.customsControlDone = true
	fmt.Println("Таможенный контроль пройден")
}

func (c *customsControl) setNext(next checkPoint) {
	c.next = next
}

func main() {
	customsControl := &customsControl{}

	regOnBoarding := &registration{}
	regOnBoarding.setNext(customsControl)

	entrance := &entranceInspection{}
	entrance.setNext(regOnBoarding)

	passenger := &passenger{name: "Alex"}
	entrance.execute(passenger)
}
