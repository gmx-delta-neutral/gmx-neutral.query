package model

import (
	"math/big"

	"github.com/google/uuid"
)

type PositionType int

const (
	Short = iota
	Long
)

type Position struct {
	Guid         uuid.UUID
	PositionType PositionType
	Coin         string
	Amount       big.Int
}
