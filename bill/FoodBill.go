package bill

import (
	"fmt"
)

// FoodBill represents the bill for food
type FoodBill struct {
	TotalAmount float32
}

// CalculateTip needs a comment so you leave me alone
func (f FoodBill) CalculateTip(b *FoodBill, t float32) (float32, error) {
	tipAmount := b.TotalAmount * (t / 100)
	fmt.Printf("Bill before tip: %v\nTip Amount at %v percent: %v\nTotal you owe: %v\n", f.TotalAmount, t, tipAmount, (tipAmount + f.TotalAmount))

	return tipAmount, nil
}
