package grpc_server

import (
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/model"
	"github.com/gmx-delta-neutral/gmx-neutral.query/pkg/api"
)

func TokenPositionDto(p *model.TokenPosition) *api.TokenPosition {
	tokenExposure := []*api.TokenExposure{}

	for _, token := range p.Exposure {
		tokenExposure = append(tokenExposure, &api.TokenExposure{
			Amount:   token.Amount.String(),
			Leverage: token.Leverage,
		})
	}

	return &api.TokenPosition{
		TokenAddress:  p.TokenAddress.String(),
		Symbol:        p.Symbol,
		Amount:        p.Amount.String(),
		Worth:         p.Worth.String(),
		Pnl:           p.PNL.String(),
		PnlPercentage: p.PNLPercentage.String(),
		CostBasis:     p.CostBasis.String(),
		Decimals:      p.Decimals,
		TokenExposure: tokenExposure,
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
		PoolAmount: p.PoolAmount.String(),
		Weight:     p.Weight.String(),
		UsdgAmount: p.UsdgAmount.String(),
		Allocation: p.Allocation.String(),
	}
}

func TransactionDto(t *model.Transaction) *api.Transaction {
	return &api.Transaction{
		TransactionId: t.TransactionId.String(),
		TokenAddress:  t.TokenAddress.String(),
		WalletAddress: t.WalletAddress.String(),
		Symbol:        t.Symbol,
		Amount:        t.Amount.String(),
		Decimals:      int32(t.Decimals),
		CostBasis:     t.CostBasis.String(),
	}
}

func TransactionDtos(transactions []*model.Transaction) []*api.Transaction {
	transactionDtos := []*api.Transaction{}

	for _, transaction := range transactions {
		transactionDtos = append(transactionDtos, TransactionDto(transaction))
	}

	return transactionDtos
}
