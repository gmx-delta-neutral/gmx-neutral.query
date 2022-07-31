package price

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	cfg "github.com/gmx-delta-neutral/gmx-neutral.query/internal/config"
	balancervault "github.com/gmx-delta-neutral/gmx-neutral.query/internal/contracts/balancer_vault"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/contracts/glp"
	glpmanager "github.com/gmx-delta-neutral/gmx-neutral.query/internal/contracts/glp_manager"
	poolstatehelper "github.com/gmx-delta-neutral/gmx-neutral.query/internal/contracts/pool_state_helper"
	pricefeed "github.com/gmx-delta-neutral/gmx-neutral.query/internal/contracts/price_feed"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/util"
)

type NextPoolState struct {
	CurrentSkew                               string `json:"currentSkew"`
	CurrentLongBalance                        string `json:"currentLongBalance"`
	CurrentLongSupply                         string `json:"currentLongSupply"`
	CurrentShortBalance                       string `json:"currentShortBalance"`
	CurrentShortSupply                        string `json:"currentShortSupply"`
	CurrentPendingLongTokenBurn               string `json:"currentPendingLongTokenBurn"`
	CurrentPendingShortTokenBurn              string `json:"currentPendingShortTokenBurn"`
	CurrentLongTokenPrice                     string `json:"currentLongTokenPrice"`
	CurrentShortTokenPrice                    string `json:"currentShortTokenPrice"`
	ExpectedSkew                              string `json:"expectedSkew"`
	ExpectedLongBalance                       string `json:"expectedLongBalance"`
	ExpectedLongSupply                        string `json:"expectedLongSupply"`
	ExpectedShortBalance                      string `json:"expectedShortBalance"`
	ExpectedShortSupply                       string `json:"expectedShortSupply"`
	ExpectedPendingLongTokenBurn              string `json:"expectedPendingLongTokenBurn"`
	ExpectedPendingShortTokenBurn             string `json:"expectedPendingShortTokenBurn"`
	ExpectedLongTokenPrice                    string `json:"expectedLongTokenPrice"`
	ExpectedShortTokenPrice                   string `json:"expectedShortTokenPrice"`
	LastOraclePrice                           string `json:"lastOraclePrice"`
	ExpectedOraclePrice                       string `json:"expectedOraclePrice"`
	ExpectedFrontRunningSkew                  string `json:"expectedFrontRunningSkew"`
	ExpectedFrontRunningLongBalance           string `json:"expectedFrontRunningLongBalance"`
	ExpectedFrontRunningLongSupply            string `json:"expectedFrontRunningLongSupply"`
	ExpectedFrontRunningShortBalance          string `json:"expectedFrontRunningShortBalance"`
	ExpectedFrontRunningShortSupply           string `json:"expectedFrontRunningShortSupply"`
	ExpectedFrontRunningPendingLongTokenBurn  string `json:"expectedFrontRunningPendingLongTokenBurn"`
	ExpectedFrontRunningPendingShortTokenBurn string `json:"expectedFrontRunningPendingShortTokenBurn"`
	TotalNetFrontRunningPendingLong           string `json:"totalNetFrontRunningPendingLong"`
	TotalNetFrontRunningPendingShort          string `json:"totalNetFrontRunningPendingShort"`
	ExpectedFrontRunningLongTokenPrice        string `json:"expectedFrontRunningLongTokenPrice"`
	ExpectedFrontRunningShortTokenPrice       string `json:"expectedFrontRunningShortTokenPrice"`
	ExpectedFrontRunningOraclePrice           string `json:"expectedFrontRunningOraclePrice"`
}

type Repository interface {
	GetPrice(priceFeedAddress common.Address) (*big.Int, error)
	GetGlpPrice() (*big.Int, error)
	GetShortBtcPrice() (*big.Int, error)
	GetTracerNextShortBtcPrice() (*big.Int, error)
}

type priceRepository struct {
	config         *cfg.Config
	arbitrumClient *ethclient.Client
}

func NewPriceRepository(config *cfg.Config, client *ethclient.Client) Repository {
	return &priceRepository{
		config:         config,
		arbitrumClient: client,
	}
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

func (r *priceRepository) GetShortBtcPrice() (*big.Int, error) {
	poolStateHelperAddress := common.HexToAddress(r.config.Addresses.PoolStateHelper)
	contract, err := poolstatehelper.NewPoolstatehelper(poolStateHelperAddress, r.arbitrumClient)

	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	leveragedPoolAddress := common.HexToAddress(r.config.Addresses.LeveragedPool)
	state, err := contract.GetExpectedState(callOpts, leveragedPoolAddress, big.NewInt(1))

	if err != nil {
		return nil, err
	}

	price := new(big.Int).Div(util.ExpandDecimals(state.ShortBalance, r.config.Decimals.Usdc), new(big.Int).Add(state.ShortSupply, state.RemainingPendingShortBurnTokens))

	return price, nil
}

func (r *priceRepository) GetBalancerShortBtcPriceFromAmount(amount *big.Int) (*big.Int, error) {
	balancerVaultAddress := common.HexToAddress(r.config.Addresses.BalancerVault)

	balancerVault, err := balancervault.NewBalancervault(balancerVaultAddress, r.arbitrumClient)

	bytes := common.Hex2Bytes("07DD9AF83C98E5AB4CBAD720E3949E1BCAAC91A60001000000000000000000A2")

	fixedBytes := [32]byte{}
	copy(fixedBytes[:], bytes)

	hash := common.BytesToHash(fixedBytes[:])
	fmt.Println(hash.String())
	usdcAddress := common.HexToAddress("0xff970a61a04b1ca14834a43f5de4533ebddb5cc8")
	shortBtcAddress := common.HexToAddress("0x6e5f70E345b4aFd271491290e026dd3d34CBb9f2")

	vaultFundManagement := balancervault.IVaultFundManagement{
		Sender:              common.HexToAddress(r.config.Addresses.Account),
		FromInternalBalance: false,
		Recipient:           common.HexToAddress(r.config.Addresses.Account),
		ToInternalBalance:   false,
	}

	swapSteps := []balancervault.IVaultBatchSwapStep{
		{
			PoolId:        fixedBytes,
			AssetInIndex:  big.NewInt(0),
			AssetOutIndex: big.NewInt(1),
			Amount:        amount,
			UserData:      big.NewInt(0).Bytes(),
		},
	}

	assets := []common.Address{shortBtcAddress, usdcAddress}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}

	amountInOut, err := balancerVault.QueryBatchSwap(callOpts, 0, swapSteps, assets, vaultFundManagement)

	if err != nil {
		return nil, err
	}

	return amountInOut[1], nil
}

func (r *priceRepository) GetTracerNextShortBtcPrice() (*big.Int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/poolsv2/nextPoolState", r.config.ThirdPartyApi.TracerApiUrl), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("poolAddress", r.config.Addresses.LeveragedPool)
	q.Add("network", "42161")

	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	nextState := &NextPoolState{}

	json.NewDecoder(res.Body).Decode(nextState)

	shortTokenPrice, success := new(big.Float).SetString(nextState.ExpectedShortTokenPrice)

	if !success {
		return nil, errors.New(fmt.Sprintf("Failed to parse short token: %s", shortTokenPrice))
	}

	expandedDecimals := new(big.Int).Exp(big.NewInt(10), big.NewInt(r.config.Decimals.Usdc), nil)
	expandedShortTokenPrice := new(big.Float).Mul(shortTokenPrice, new(big.Float).SetInt(expandedDecimals))
	fmt.Println(expandedShortTokenPrice)
	truncatedShortTokenPrice := new(big.Int)
	expandedShortTokenPrice.Int(truncatedShortTokenPrice)
	return truncatedShortTokenPrice, nil
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

	expandedGlpPrice := new(big.Int).Div(util.ExpandDecimals(aum, r.config.Decimals.Usdc), supply)
	return expandedGlpPrice, nil
}
