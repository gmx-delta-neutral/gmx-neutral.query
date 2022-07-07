package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TokenPosition struct {
	TokenAddress  common.Address
	WalletAddress common.Address
	Symbol        string
	Amount        *big.Int
	Worth         *big.Int
	PNL           *big.Int
}

type ShortPosition struct {
	TokenAddress  common.Address
	WalletAddress common.Address
	Symbol        string
	Balance       big.Int
	Decimals      int32
	PNL           big.Int
}
