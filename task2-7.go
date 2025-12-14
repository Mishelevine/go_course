package main

import (
	"fmt"
	"time"
)

type Transport struct {
	model            string
	manufacture_year int
	mileage          float32
}

type Car struct {
	Transport
	avg_fuel_consuption float32
}

func (t *Transport) SendOnTrip(distance float32) {
	t.mileage += distance
}

func (t Transport) String() string {
	return fmt.Sprintf("Модель %s %d г.в., пробег %.2f км", t.model, t.manufacture_year, t.mileage)
}

func (c Car) String() string {
	return fmt.Sprintf("Автомобиль: %s, расход топлива %.1f л/км", c.Transport, c.avg_fuel_consuption)
}

func NewTransport(model string) Transport {
	new_tr := Transport{model: model, manufacture_year: time.Now().Year(), mileage: 0}
	return new_tr
}

func NewAuto(model string, fCons float32) Car {
	new_car := Car{Transport: NewTransport(model), avg_fuel_consuption: fCons}
	return new_car
}

func task2_7() {
	a1 := NewAuto("Geely Coolray", 8.5)
	t1 := NewTransport("самолет МС-21")
	fmt.Println(a1)
	fmt.Println(t1)
	a1.SendOnTrip(100)
	t1.SendOnTrip(1500)
	fmt.Println(a1)
	fmt.Println(t1)
}
