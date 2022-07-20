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
	config             *cfg.Config
	tokenRepository    Repository
	priceRepository    price.Repository
	transactionService transaction.Service
}

func NewService(config *cfg.Config, tokenRepository Repository, priceRepository price.Repository, transactionService transaction.Service) Service {
	return &tokenService{
		config:             config,
		tokenRepository:    tokenRepository,
		priceRepository:    priceRepository,
		transactionService: transactionService,
	}
}

var glpToken = common.HexToAddress("0x01234181085565ed162a948b6a5e88758CD7c7b8")

type TokenPositionHandler = func(common.Address) (*model.TokenPosition, error)

var tokenPositionHandlerMappings = map[common.Address]func(*tokenService) (*model.TokenPosition, error){
	glpToken: func(service *tokenService) (*model.TokenPosition, error) { return service.getGlpPosition(glpToken) },
}

func (service *tokenService) GetTokenPosition(tokenAddress common.Address) (*model.TokenPosition, error) {
	if tokenPositionHandler, ok := tokenPositionHandlerMappings[tokenAddress]; ok {
		return tokenPositionHandler(service)
	}

	return nil, errors.New("Token mapping not foundd")
}

func (service *tokenService) getGlpPosition(tokenAddress common.Address) (*model.TokenPosition, error) {
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

	for _, transaction := range *transactions {
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
		TokenAddress:  tokenAddress,
		WalletAddress: tokenAddress,
		Symbol:        "BTC",
		Amount:        amount,
		Worth:         worth,
		CostBasis:     totalCostBasis,
		PNL:           pnl,
		PNLPercentage: new(big.Float).Quo(new(big.Float).SetInt(pnl), new(big.Float).SetInt(totalCostBasis)),
	}

	return position, err

}

func (service *tokenService) GetTokenPositions() ([]*model.TokenPosition, error) {
	tokenPositions := []*model.TokenPosition{}

	for _, handler := range tokenPositionHandlerMappings {
		tokenPosition, err := handler(service)
		if err != nil {
			return nil, err
		}
		tokenPositions = append(tokenPositions, tokenPosition)
	}

	return tokenPositions, nil
}
