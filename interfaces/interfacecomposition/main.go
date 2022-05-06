package main

import "fmt"

type Payer interface {
	Pay()
}

type Ringer interface {
	Ring()
	Call()
}

type NFCPhone interface {
	Payer
	Ringer
}

type Phone struct {
	Model string
}

func MakeAllOperations(phone NFCPhone) {
	phone.Pay()
	phone.Ring()
	phone.Call()
}

func (*Phone) Call() {
	fmt.Println("Звоню кому-то!")
}

func (*Phone) Ring() {
	fmt.Println("Звонит кто-то!")
}

func (*Phone) Pay() {
	fmt.Println("Плачу за что-то!")
}

func main() {
	phone := Phone{"123"}
	MakeAllOperations(&phone)
}
