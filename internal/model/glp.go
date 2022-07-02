package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type GlpAsset struct {
	TokenAddress common.Address
	Symbol       string
	PoolAmount   big.Int
	Weight       big.Int
	UsdgAmount   big.Int
	Allocation   big.Int
}
