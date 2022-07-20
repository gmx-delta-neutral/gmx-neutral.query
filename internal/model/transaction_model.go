package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	TransactionId common.Hash
	TokenAddress  common.Address
	WalletAddress common.Address
	Symbol        string
	Amount        *big.Int
	Decimals      int
	CostBasis     *big.Int
}
