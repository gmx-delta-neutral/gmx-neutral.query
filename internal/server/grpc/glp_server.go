package grpc_server

import (
	"context"

	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/glp"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/pkg/api"
)

func NewGlpServer(glpService glp.Service) *GlpServer {
	return &GlpServer{
		glpService: glpService,
	}
}

type GlpServer struct {
	glpService glp.Service
}

func (p *GlpServer) GetGlpAssets(ctx context.Context, request *api.GetGlpAssetsRequest) (*api.GetGlpAssetsResponse, error) {
	glpAssetsResponse := []*api.GlpAsset{}
	res := api.GetGlpAssetsResponse{
		Assets: glpAssetsResponse,
	}

	glp_assets, err := p.glpService.GetGlpAssets()

	if err != nil {
		return &res, err
	}

	for _, glp_asset := range glp_assets {
		glpAssetsResponse = append(glpAssetsResponse, GlpAssetDto(glp_asset))
	}

	res.Assets = glpAssetsResponse
	return &res, nil
}
