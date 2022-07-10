package token

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/model"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/price"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/util"
)

type Service interface {
	GetTokenPositions() ([]*model.TokenPosition, error)
	GetTokenPosition(tokenAddress common.Address) (*model.TokenPosition, error)
}

type tokenService struct {
	tokenRepository Repository
	priceRepository price.Repository
}

func NewService(tokenRepository Repository, priceRepository price.Repository) Service {
	return &tokenService{tokenRepository: tokenRepository, priceRepository: priceRepository}
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
	transactions, err := service.tokenRepository.GetTokenTransactions(tokenAddress)

	if err != nil {
		return nil, err
	}

	amount := new(big.Int)
	worth := new(big.Int)
	pnl := new(big.Int)

	for _, transaction := range transactions {
		transactionWorth := util.RemoveDecimals(new(big.Int).Mul(transaction.Amount, glpPrice), 18)
		purchaseWorth := util.RemoveDecimals(new(big.Int).Mul(transaction.Amount, transaction.PurchasePrice), 18)
		transactionPnl := new(big.Int).Sub(transactionWorth, purchaseWorth)
		amount = new(big.Int).Add(amount, transaction.Amount)
		worth = new(big.Int).Add(worth, transactionWorth)
		pnl = new(big.Int).Add(pnl, transactionPnl)
	}

	position := &model.TokenPosition{
		TokenAddress:  tokenAddress,
		WalletAddress: tokenAddress,
		Symbol:        "BTC",
		Amount:        amount,
		Worth:         worth,
		PNL:           pnl,
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
