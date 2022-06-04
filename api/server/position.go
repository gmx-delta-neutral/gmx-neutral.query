package server

import (
	"context"

	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/api/generated"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/position"
	"github.com/google/uuid"
)

func NewPositionServer(positionService position.Service) *PositionServer {
	return &PositionServer{
		positionService: positionService,
	}
}

type PositionServer struct {
	positionService position.Service
}

func (p *PositionServer) GetPositions(ctx context.Context, request *generated.GetPositionsRequest) (*generated.GetPositionsResponse, error) {
	positionsResponse := []*generated.Position{}
	res := generated.GetPositionsResponse{
		Positions: positionsResponse,
	}

	positions, err := p.positionService.GetPositions()

	if err != nil {
		return &res, err
	}

	for _, position := range positions {
		positionsResponse = append(positionsResponse, PositionDto(position))
	}

	res.Positions = positionsResponse
	return &res, nil
}

func (p *PositionServer) GetPosition(ctx context.Context, request *generated.GetPositionRequest) (*generated.GetPositionResponse, error) {
	res := generated.GetPositionResponse{
		Position: &generated.Position{},
	}

	position, err := p.positionService.GetPosition(uuid.MustParse(request.Guid))

	if err != nil {
		return &res, err
	}

	res.Position = PositionDto(position)
	return &res, nil
}
