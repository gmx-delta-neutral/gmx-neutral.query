package grpc_server

import (
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/model"
	"github.com/gmx-delta-neutral/gmx-neutral.query/pkg/api"
)

func TokenPositionDto(p *model.TokenPosition) *api.TokenPosition {
	return &api.TokenPosition{
		TokenAddress:  p.TokenAddress.String(),
		Symbol:        p.Symbol,
		Amount:        p.Amount.Bytes(),
		Worth:         p.Worth.Bytes(),
		Pnl:           p.PNL.Bytes(),
		PnlPercentage: p.PNLPercentage.String(),
	}
}

func TokenPositionDtos(tokenPositions []*model.TokenPosition) []*api.TokenPosition {
	tokenPositionDtos := []*api.TokenPosition{}

	for _, tokenPosition := range tokenPositions {
		tokenPositionDtos = append(tokenPositionDtos, TokenPositionDto(tokenPosition))
	}

	return tokenPositionDtos
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
