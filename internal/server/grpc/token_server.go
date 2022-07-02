package grpc_server

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/token"
	"github.com/gmx-delta-neutral/gmx-neutral.query/pkg/api"
)

func NewTokenServer(tokenService token.Service) *TokenServer {
	return &TokenServer{
		tokenService: tokenService,
	}
}

type TokenServer struct {
	tokenService token.Service
}

func (p *TokenServer) GetTokenPosition(ctx context.Context, request *api.GetTokenPositionRequest) (*api.GetTokenPositionResponse, error) {
	tokenAddress := common.HexToAddress(request.TokenAddress)
	tokenPosition, err := p.tokenService.GetTokenPosition(tokenAddress)

	if err != nil {
		return nil, err
	}

	tokenPositionResponse := api.GetTokenPositionResponse{
		TokenPosition: TokenPositionDto(tokenPosition),
	}
	return &tokenPositionResponse, nil
}
