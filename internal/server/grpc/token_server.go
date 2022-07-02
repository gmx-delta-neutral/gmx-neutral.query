package grpc_server

import (
	"context"

	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/api/generated"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/token"
	"github.com/ethereum/go-ethereum/common"
)

func NewTokenServer(tokenService token.Service) *TokenServer {
	return &TokenServer{
		tokenService: tokenService,
	}
}

type TokenServer struct {
	tokenService token.Service
}

func (p *TokenServer) GetTokenPosition(ctx context.Context, request *generated.GetTokenPositionRequest) (*generated.GetTokenPositionResponse, error) {
	tokenAddress := common.HexToAddress(request.TokenAddress)
	tokenPosition, err := p.tokenService.GetTokenPosition(tokenAddress)

	if err != nil {
		return nil, err
	}

	tokenPositionResponse := generated.GetTokenPositionResponse{
		TokenPosition: TokenPositionDto(tokenPosition),
	}
	return &tokenPositionResponse, nil
}
