package grpc_server

import (
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/model"
	"github.com/gmx-delta-neutral/gmx-neutral.query/pkg/api"
)

func TokenPositionDto(p model.TokenPosition) *api.TokenPosition {
	return &api.TokenPosition{
		TokenAddress: p.TokenAddress.String(),
		Symbol:       p.Symbol,
		Balance:      p.Balance.Bytes(),
		Decimals:     p.Decimals,
	}
}

func GlpAssetDto(p model.GlpAsset) *api.GlpAsset {
	return &api.GlpAsset{
		Symbol:     p.Symbol,
		PoolAmount: p.PoolAmount.Bytes(),
		Weight:     p.Weight.Bytes(),
		UsdgAmount: p.UsdgAmount.Bytes(),
		Allocation: p.Allocation.Bytes(),
	}
}
