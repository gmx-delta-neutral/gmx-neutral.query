package server

import (
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/api/generated"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/model"
)

func PositionDto(p model.Position) *generated.Position {
	return &generated.Position{
		Guid:         p.Guid.String(),
		PositionType: generated.PositionType(p.PositionType),
		Coin:         p.Coin,
		Amount:       p.Amount.Bytes(),
	}
}
