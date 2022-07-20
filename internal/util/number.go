package util

import "math/big"

func ExpandDecimals(n *big.Int, decimals int64) *big.Int {
	return new(big.Int).Mul(n, new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals), nil))
}

// RemoveDecimals does this
func RemoveDecimals(n *big.Int, decimals int64) *big.Int {
	return new(big.Int).Div(n, new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals), nil))
}
