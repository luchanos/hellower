package main

import (
	"fmt"
)

type Machinegun struct {
	Ammo int
}

type Car struct {
	Fuel         int
	EngineStatus string
}

type WarCar struct {
	*Machinegun
	*Car
}

func (wc *WarCar) OpenFire() {
	if wc.Ammo > 0 {
		fmt.Println("open fire by", wc)
	} else {
		fmt.Println("not enough ammo")
	}
}

func (wc *WarCar) Move() {
	fmt.Println("wc is moving")
}

func (mg *Machinegun) OpenFire() {
	if mg.Ammo > 0 {
		fmt.Println("open fire by", mg)
	} else {
		fmt.Println("not enough ammo")
	}
}

func (car *Car) Move() {
	fmt.Println("moving")
}

type WarTransport interface {
	OpenFire()
	Move()
}

func MakeWarcAction(wt WarTransport) {
	wt.OpenFire()
	wt.Move()
}

func main() {
	wc := WarCar{&Machinegun{1}, &Car{Fuel: 1, EngineStatus: "good"}}
	MakeWarcAction(&wc)
}
