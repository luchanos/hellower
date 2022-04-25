package chapter1

import "fmt"

func three() {
	price := 100
	fmt.Println("Price is", price, "dollars")

	var taxRate float64 = 0.08
	tax := float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars")

	total := float64(price) + tax
	fmt.Println("Total cost is", total, "dollars")

	availableFunds := 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", total <= float64(availableFunds))
}
