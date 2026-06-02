package operations

import (
	"errors"

	"github.com/shopspring/decimal"
)

func Add(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) {

	if a.IsNegative() || b.IsNegative() {
		return decimal.Zero, errors.New("inputs must be positive")
	}

	return a.Add(b), nil
}
