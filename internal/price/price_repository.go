package price

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/contracts/glp"
	glpmanager "github.com/gmx-delta-neutral/gmx-neutral.query/internal/contracts/glp_manager"
	pricefeed "github.com/gmx-delta-neutral/gmx-neutral.query/internal/contracts/price_feed"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/util"
)

type Repository interface {
	GetPrice(priceFeedAddress common.Address) (*big.Int, error)
	GetGlpPrice() (*big.Int, error)
}

type priceRepository struct {
}

func NewPriceRepository() Repository {
	return &priceRepository{}
}

var glpManagerAddress = common.HexToAddress("0xe1ae4d4b06A5Fe1fc288f6B4CD72f9F8323B107F")
var glpAddress = common.HexToAddress("0x01234181085565ed162a948b6a5e88758CD7c7b8")

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

func (r *priceRepository) GetGlpPrice() (*big.Int, error) {
	client, err := ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")

	if err != nil {
		return nil, err
	}

	ctr, err := glpmanager.NewGlpmanager(glpManagerAddress, client)

	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	aum, err := ctr.GetAumInUsdg(callOpts, true)

	if err != nil {
		return nil, err
	}

	glpContract, err := glp.NewGlp(glpAddress, client)

	if err != nil {
		return nil, err
	}

	supply, err := glpContract.TotalSupply(callOpts)

	if err != nil {
		return nil, err
	}

	expandedGlpPrice := new(big.Int).Div(util.ExpandDecimals(aum, 18), supply)
	return expandedGlpPrice, nil
}
