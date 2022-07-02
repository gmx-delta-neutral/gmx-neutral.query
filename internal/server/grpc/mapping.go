package grpc_server

import (
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/api/generated"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/model"
)

func TokenPositionDto(p model.TokenPosition) *generated.TokenPosition {
	return &generated.TokenPosition{
		TokenAddress: p.TokenAddress.String(),
		Symbol:       p.Symbol,
		Balance:      p.Balance.Bytes(),
		Decimals:     p.Decimals,
	}
}

func GlpAssetDto(p model.GlpAsset) *generated.GlpAsset {
	return &generated.GlpAsset{
		Symbol:     p.Symbol,
		PoolAmount: p.PoolAmount.Bytes(),
		Weight:     p.Weight.Bytes(),
		UsdgAmount: p.UsdgAmount.Bytes(),
		Allocation: p.Allocation.Bytes(),
	}
}
