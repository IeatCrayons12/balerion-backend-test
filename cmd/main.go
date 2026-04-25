package main

import (
	"fmt"

	"github.com/IeatCrayons12/balerion-backend-test/internal/thaibaht"
	"github.com/shopspring/decimal"
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}
	for _, input := range inputs {
		fmt.Println(input)
		fmt.Println(thaibaht.Convert(input))
		fmt.Println()
	}
}
