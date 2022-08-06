package token

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	cfg "github.com/gmx-delta-neutral/gmx-neutral.query/internal/config"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/model"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/price"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/transaction"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/util"
)

type Service interface {
	GetTokenPositions() ([]*model.TokenPosition, error)
	GetTokenPosition(tokenAddress common.Address) (*model.TokenPosition, error)
}

type tokenService struct {
	config                       *cfg.Config
	tokenRepository              Repository
	priceRepository              price.Repository
	transactionService           transaction.Service
	tokenPositionHandlerMappings map[string]func(*tokenService) (*model.TokenPosition, error)
}

func NewService(config *cfg.Config, tokenRepository Repository, priceRepository price.Repository, transactionService transaction.Service) Service {

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

	fmt.Printf("Cost basis:  %s\n", totalCostBasis)
	fmt.Printf("Amount:  %s\n", amount)
	fmt.Printf("Glp price: %s\n", glpPrice)

	worth := util.RemoveDecimals(new(big.Int).Mul(amount, glpPrice), service.config.Decimals.Glp)
	fmt.Printf("How much it is worth now: %s\n", worth)
	fmt.Printf("What is the cost basis: %s\n", totalCostBasis)
	pnl := new(big.Int).Sub(worth, totalCostBasis)

	fmt.Printf("Worth: %s\n", worth)
	fmt.Printf("Pnl: %s\n", pnl)

	position := &model.TokenPosition{
		TokenAddress:  common.HexToAddress(service.config.Addresses.ShortBtcToken),
		WalletAddress: walletAddress,
		Symbol:        "BTC",
		Amount:        amount,
		Worth:         worth,
		CostBasis:     totalCostBasis,
		PNL:           pnl,
		PNLPercentage: new(big.Float).Quo(new(big.Float).SetInt(pnl), new(big.Float).SetInt(totalCostBasis)),
		Decimals:      int32(service.config.Decimals.Glp),
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

	fmt.Printf("How much it is worth now: %s\n", worth)
	fmt.Printf("What is the cost basis: %s\n", totalCostBasis)
	pnl := new(big.Int).Sub(worth, totalCostBasis)

	fmt.Printf("Worth: %s\n", worth)
	fmt.Printf("Pnl: %s\n", pnl)

	position := &model.TokenPosition{
		TokenAddress:  common.HexToAddress(service.config.Addresses.ShortBtcToken),
		WalletAddress: common.HexToAddress(service.config.Addresses.Account),
		Symbol:        "ShortBtc3X",
		Amount:        amount,
		Worth:         worth,
		CostBasis:     totalCostBasis,
		PNL:           pnl,
		PNLPercentage: new(big.Float).Quo(new(big.Float).SetInt(pnl), new(big.Float).SetInt(totalCostBasis)),
		Decimals:      int32(service.config.Id.TracerThreeLeverageShortBtc),
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
