package token

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	cfg "github.com/gmx-delta-neutral/gmx-neutral.query/internal/config"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/glp"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/model"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/price"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/transaction"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/util"
	"golang.org/x/exp/slices"
)

type Service interface {
	GetTokenPositions() ([]*model.TokenPosition, error)
	GetTokenPosition(tokenAddress common.Address) (*model.TokenPosition, error)
}

type tokenService struct {
	config                       *cfg.Config
	tokenRepository              Repository
	priceRepository              price.Repository
	glpService                   glp.Service
	transactionService           transaction.Service
	tokenPositionHandlerMappings map[string]func(*tokenService) (*model.TokenPosition, error)
}

func NewService(config *cfg.Config, tokenRepository Repository, priceRepository price.Repository, transactionService transaction.Service, glpService glp.Service) Service {
	var tokenPositionHandlerMappings = map[string]func(*tokenService) (*model.TokenPosition, error){
		config.Addresses.Glp: func(service *tokenService) (*model.TokenPosition, error) {
			return service.getGlpPosition()
		},
		config.Addresses.ShortBtcToken: func(service *tokenService) (*model.TokenPosition, error) {
			return service.getShortBtcPosition()
		},
	}

	return &tokenService{
		config:                       config,
		tokenRepository:              tokenRepository,
		priceRepository:              priceRepository,
		transactionService:           transactionService,
		tokenPositionHandlerMappings: tokenPositionHandlerMappings,
		glpService:                   glpService,
	}
}

type TokenPositionHandler = func(common.Address) (*model.TokenPosition, error)

func (service *tokenService) GetTokenPosition(tokenAddress common.Address) (*model.TokenPosition, error) {
	if tokenPositionHandler, ok := service.tokenPositionHandlerMappings[tokenAddress.String()]; ok {
		return tokenPositionHandler(service)
	}

	return nil, errors.New("token mapping not foundd")
}

func (service *tokenService) getGlpPosition() (*model.TokenPosition, error) {
	glpPrice, err := service.priceRepository.GetGlpPrice()
	if err != nil {
		return nil, err
	}

	transactions, err := service.transactionService.GetGlpTransactions()

	if err != nil {
		return nil, err
	}

	amount := new(big.Int)
	totalCostBasis := new(big.Int)

	for _, transaction := range transactions {
		amount = new(big.Int).Add(amount, transaction.Amount)
		totalCostBasis = new(big.Int).Add(totalCostBasis, transaction.CostBasis)
	}

	worth := util.RemoveDecimals(new(big.Int).Mul(amount, glpPrice), service.config.Decimals.Glp)
	pnl := new(big.Int).Sub(worth, totalCostBasis)

	glpExposure, err := service.glpService.GetGlpExposure(worth)

	if err != nil {
		return nil, err
	}

	exposureTokens := []string{"WBTC", "WETH"}

	filteredGlpExposure := []*model.TokenExposure{}

	for _, exposure := range glpExposure {
		if slices.Contains(exposureTokens, exposure.Symbol) {
			filteredGlpExposure = append(filteredGlpExposure, exposure)
		}
	}

	position := &model.TokenPosition{
		TokenAddress:  common.HexToAddress(service.config.Addresses.Glp),
		WalletAddress: walletAddress,
		Symbol:        "GLP",
		Amount:        amount,
		Worth:         worth,
		CostBasis:     totalCostBasis,
		PNL:           pnl,
		PNLPercentage: new(big.Float).Quo(new(big.Float).SetInt(pnl), new(big.Float).SetInt(totalCostBasis)),
		Decimals:      int32(service.config.Decimals.Glp),
		Exposure:      filteredGlpExposure,
	}

	return position, err
}

func (service *tokenService) getShortBtcPosition() (*model.TokenPosition, error) {
	transactions, err := service.transactionService.GetTracerTransactions()

	if err != nil {
		return nil, err
	}

	amount := new(big.Int)
	totalCostBasis := new(big.Int)

	for _, transaction := range transactions {
		amount = new(big.Int).Add(amount, transaction.Amount)
		totalCostBasis = new(big.Int).Add(totalCostBasis, transaction.CostBasis)
	}

	price, err := service.priceRepository.GetShortBtcPrice()
	worth := util.RemoveDecimals(new(big.Int).Mul(price, amount), service.config.Decimals.Usdc)

	if err != nil {
		return nil, err
	}

	pnl := new(big.Int).Sub(worth, totalCostBasis)

	position := &model.TokenPosition{
		TokenAddress:  common.HexToAddress(service.config.Addresses.ShortBtcToken),
		WalletAddress: common.HexToAddress(service.config.Addresses.Account),
		Symbol:        "ShortBtc3X",
		Amount:        amount,
		Worth:         worth,
		CostBasis:     totalCostBasis,
		PNL:           pnl,
		PNLPercentage: new(big.Float).Quo(new(big.Float).SetInt(pnl), new(big.Float).SetInt(totalCostBasis)),
		Decimals:      int32(service.config.Decimals.ShortBtc3X),
		Exposure: []*model.TokenExposure{
			{
				Amount:   new(big.Int).Mul(amount, big.NewInt(3)),
				Leverage: 3,
				Symbol:   "ShortBtc3X",
			},
		},
	}

	return position, err

}

func (service *tokenService) GetTokenPositions() ([]*model.TokenPosition, error) {
	tokenPositions := []*model.TokenPosition{}

	for _, handler := range service.tokenPositionHandlerMappings {
		tokenPosition, err := handler(service)
		if err != nil {
			return nil, err
		}
		tokenPositions = append(tokenPositions, tokenPosition)
	}

	return tokenPositions, nil
}
