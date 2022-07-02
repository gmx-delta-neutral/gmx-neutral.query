package price

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	pricefeed "github.com/gmx-delta-neutral/gmx-neutral.query/internal/contracts/price_feed"
)

type Repository interface {
	GetPrice(priceFeedAddress common.Address) (*big.Int, error)
}

type priceRepository struct {
}

func NewPriceRepository() Repository {
	return &priceRepository{}
}

func (r *priceRepository) GetPrice(priceFeedAddress common.Address) (*big.Int, error) {
	client, err := ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")

	if err != nil {
		return nil, err
	}

	ctr, err := pricefeed.NewPricefeed(priceFeedAddress, client)

	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	result, err := ctr.LatestRoundData(callOpts)

	if err != nil {
		return nil, err
	}

	return result.Answer, nil
}
