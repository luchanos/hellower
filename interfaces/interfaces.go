package main

import "fmt"

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}

type Card struct {
	Balance        int
	ValidUntil     string
	CardholderName string
	CVV            int
	Number         string
}

type ApplePay struct {
	Money   int
	AppleId string
}

func (w *ApplePay) Pay(n int) error {
	if w.Money < n {
		return fmt.Errorf("not enough money")
	}
	w.Money -= n
	fmt.Println("спасибо за покупку эпл пэем!")
	return nil
}

func (w *Card) Pay(n int) error {
	if w.Balance < n {
		return fmt.Errorf("not enough money")
	}
	w.Balance -= n
	fmt.Println("спасибо за покупку карточкой!")
	return nil
}

func (w *Wallet) Pay(n int) error {
	if w.Cash < n {
		return fmt.Errorf("not enough money")
	}
	w.Cash -= n
	fmt.Println("спасибо за покупку!")
	return nil
}

func Buy(p Payer, v int) {
	err := p.Pay(v)
	if err != nil {
		fmt.Println(err)
	}
}

func Buy2(p Payer, v int) {
	switch p.(type) {
	case *Wallet:
		fmt.Println("Оплата наличными")
	case *Card:
		plasticCard, ok := p.(*Card)
		if !ok {
			panic("не удалось преобразовать к типу Card")
		}
		fmt.Println("оплата картой", plasticCard)
	default:
		fmt.Println("что-то новое")
	}
}

func main() {
	wallet := &Wallet{100}
	Buy(wallet, 10)

	var MyMoney Payer
	MyMoney = &ApplePay{100, "100"}
	Buy(MyMoney, 100)

	MyMoney = &Card{100, "2022", "Nik", 333, "23321"}
	Buy(MyMoney, 101)
}
