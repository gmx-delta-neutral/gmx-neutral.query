package grpc_server

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/price"
	"github.com/gmx-delta-neutral/gmx-neutral.query/pkg/api"
)

func NewPriceServer(priceService price.Service) api.PriceServiceServer {
	return &PriceServer{
		priceService: priceService,
	}
}

type PriceServer struct {
	priceService price.Service
}

func (s *PriceServer) GetTokenPrice(ctx context.Context, request *api.GetTokenPriceRequest) (*api.GetTokenPriceResponse, error) {
	address := common.BytesToAddress(request.Address)
	price, err := s.priceService.GetPrice(address)

	if err != nil {
		return nil, err
	}

	response := &api.GetTokenPriceResponse{
		Price: price.Bytes(),
	}

	return response, nil
}
