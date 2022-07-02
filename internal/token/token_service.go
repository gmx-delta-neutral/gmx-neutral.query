package token

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/glp"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/model"
)

type Service interface {
	GetTokenPositions() ([]*model.TokenPosition, error)
	GetTokenPosition(tokenAddress common.Address) (*model.TokenPosition, error)
}

type tokenService struct {
	tokenRepository Repository
	glpRepository   glp.Repository
}

func NewService(tokenRepository Repository) Service {
	return &tokenService{tokenRepository: tokenRepository}
}

func (service *tokenService) GetTokenPosition(tokenAddress common.Address) (*model.TokenPosition, error) {
	position, err := service.tokenRepository.GetTokenPosition(tokenAddress)

	if err != nil {
		return nil, err
	}

	return position, err
}

func (service *tokenService) GetTokenPositions() ([]*model.TokenPosition, error) {
	position, err := service.tokenRepository.GetTokenPositions()

	if err != nil {
		return nil, err
	}

	return position, err
}
