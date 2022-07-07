package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type Transaction struct {
	TransactionId uuid.UUID
	TokenAddress  common.Address
	WalletAddress common.Address
	Symbol        string
	Amount        *big.Int
	Decimals      int
	PurchasePrice *big.Int
}
