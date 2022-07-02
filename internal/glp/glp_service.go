package glp

import (
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/model"
)

type Service interface {
	GetGlpAssets() ([]model.GlpAsset, error)
}

type service struct {
	glpRepository Repository
}

func NewService(glpRepository Repository) *service {
	return &service{glpRepository: glpRepository}
}

func (service *service) GetGlpAssets() ([]model.GlpAsset, error) {
	positions, err := service.glpRepository.GetGlpAssets()

	if err != nil {
		return nil, err
	}

	return positions, nil
}
