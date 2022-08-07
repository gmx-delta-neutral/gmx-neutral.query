package glp

import (
	"math/big"

	cfg "github.com/gmx-delta-neutral/gmx-neutral.query/internal/config"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/model"
)

type Service interface {
	GetGlpAssets() ([]model.GlpAsset, error)
	GetGlpExposure(glpWorth *big.Int) ([]*model.TokenExposure, error)
}

type service struct {
	config        *cfg.Config
	glpRepository Repository
}

func NewService(glpRepository Repository, config *cfg.Config) *service {
	return &service{
		config:        config,
		glpRepository: glpRepository,
	}
}

func (service *service) GetGlpAssets() ([]model.GlpAsset, error) {
	positions, err := service.glpRepository.GetGlpAssets()

	if err != nil {
		return nil, err
	}

	return positions, nil
}

func (service *service) GetGlpExposure(glpWorth *big.Int) ([]*model.TokenExposure, error) {
	tokenExposure := []*model.TokenExposure{}

	assets, err := service.GetGlpAssets()

	if err != nil {
		return nil, err
	}

	for _, asset := range assets {
		amount := new(big.Int).Div(new(big.Int).Mul(glpWorth, &asset.Allocation), big.NewInt(service.config.Decimals.PercentDivisor))
		tokenExposure = append(tokenExposure, &model.TokenExposure{
			Amount:   amount,
			Leverage: 1,
			Symbol:   asset.Symbol,
		})
	}

	return tokenExposure, nil
}
