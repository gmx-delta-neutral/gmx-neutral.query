package grpc_server

import (
	"context"

	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/api/generated"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/price"
	"github.com/ethereum/go-ethereum/common"
)

func NewPriceServer(priceService price.Service) generated.PriceServiceServer {
	return &PriceServer{
		priceService: priceService,
	}
}

type PriceServer struct {
	priceService price.Service
}

func (s *PriceServer) GetTokenPrice(ctx context.Context, request *generated.GetTokenPriceRequest) (*generated.GetTokenPriceResponse, error) {
	address := common.BytesToAddress(request.Address)
	price, err := s.priceService.GetPrice(address)

	if err != nil {
		return nil, err
	}

	response := &generated.GetTokenPriceResponse{
		Price: price.Bytes(),
	}

	return response, nil
}
