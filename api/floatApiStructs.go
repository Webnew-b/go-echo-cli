package api

import (
	"fmt"
)

type Float64TwoPrecision float64

func (f *Float64TwoPrecision) MarshalJSON() ([]byte, error) {
	num := float64(*f)
	return []byte(fmt.Sprintf("%.2f", num)), nil
}
