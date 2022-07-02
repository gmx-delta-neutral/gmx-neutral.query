package token

import (
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/glp"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/model"
	"github.com/ethereum/go-ethereum/common"
)

type Service interface {
	GetTokenPosition(tokenAddress common.Address) (model.TokenPosition, error)
}

type tokenService struct {
	tokenRepository Repository
	glpRepository   glp.Repository
}

func NewService(tokenRepository Repository) Service {
	return &tokenService{tokenRepository: tokenRepository}
}

func (service *tokenService) GetTokenPosition(tokenAddress common.Address) (model.TokenPosition, error) {
	position, err := service.tokenRepository.GetTokenPosition(tokenAddress)

	if err != nil {
		return position, err
	}

	// glp_assets, err := service.glpRepository.GetGlpAssets()

	// if err != nil {
	// 	return position, err
	// }

	// current_price := big.NewInt(0)
	// for _, glp_asset := range glp_assets {
	// 	if glp_asset.TokenAddress == tokenAddress {
	// 		current_price := glp_asset.
	// 	}

	// }
	// position.PNL =

	return position, err
}
