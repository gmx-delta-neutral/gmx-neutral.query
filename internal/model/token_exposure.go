package model

import "math/big"

type TokenExposure struct {
	Amount   *big.Int
	Leverage float64
	Symbol   string
}
