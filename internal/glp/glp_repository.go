package glp

import (
	"context"
	"log"
	"math/big"

	vaultreader "github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/contracts/vault_reader"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient"
)

const vaultPropsLength = 14

type Repository interface {
	GetGlpAssets() ([]model.GlpAsset, error)
}
type repository struct{}

func NewRepository() Repository {
	return repository{}
}

func extractGlpAssetsFromBinary(rawInfo []*big.Int, glp_assets []GlpAssetDefinition) []model.GlpAsset {
	total_supply := new(big.Int)

	for i := 0; i < len(glp_assets); i++ {
		usdg_amount := *rawInfo[i*vaultPropsLength+2]
		total_supply = new(big.Int).Add(total_supply, &usdg_amount)
	}

	tokenInfos := []model.GlpAsset{}
	for i := 0; i < len(glp_assets); i++ {
		pool_amount := rawInfo[i*vaultPropsLength]
		usdg_amount := *rawInfo[i*vaultPropsLength+2]

		allocation := new(big.Int).Div(new(big.Int).Mul(&usdg_amount, big.NewInt(10000)), total_supply)

		vaultTokenInfo := model.GlpAsset{
			TokenAddress: common.HexToAddress(glp_assets[i].Address),
			Symbol:       glp_assets[i].Symbol,
			PoolAmount:   *pool_amount,
			UsdgAmount:   *rawInfo[i*vaultPropsLength+2],
			Weight:       *rawInfo[i*vaultPropsLength+4],
			Allocation:   *allocation,
		}

		tokenInfos = append(tokenInfos, vaultTokenInfo)
	}

	return tokenInfos
}

type GlpAssetDefinition struct {
	Symbol  string
	Address string
}

var avax_asset_definitions = []GlpAssetDefinition{
	{
		Symbol:  "WAVAX",
		Address: "0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7",
	},
	{
		Symbol:  "WBTC",
		Address: "0x50b7545627a5162f82a992c33b87adc75187b218",
	},
	{
		Symbol:  "WETH",
		Address: "0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab",
	},
	{
		Symbol:  "USDC",
		Address: "0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e",
	},
	{
		Symbol:  "USDC.e",
		Address: "0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664",
	},
}

func (repository repository) GetGlpAssets() ([]model.GlpAsset, error) {
	vaultReaderAddress := common.HexToAddress("0x80785f96743d5Aef7725d88256fdBCfF43fBd112")
	vaultAddress := common.HexToAddress("0x9ab2De34A33fB459b538c43f251eB825645e8595")
	positionManagerAddress := common.HexToAddress("0xF2ec2e52c3b5F8b8bd5A3f93945d05628A233216")
	avaxAddress := common.HexToAddress("0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7")

	tokens := []common.Address{}
	for _, token := range avax_asset_definitions {
		tokens = append(tokens, common.HexToAddress(token.Address))
	}

	client, err := ethclient.Dial("https://api.avax.network/ext/bc/C/rpc")

	if err != nil {
		return nil, err
	}

	ctr, err := vaultreader.NewVaultreader(vaultReaderAddress, client)

	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	result, err := ctr.GetVaultTokenInfoV3(callOpts, vaultAddress, positionManagerAddress, avaxAddress, math.BigPow(10, 18), tokens)

	if err != nil {
		log.Fatalf("Error getting vault token info. Error: %s", err)
	}

	glp_assets := extractGlpAssetsFromBinary(result, avax_asset_definitions)

	return glp_assets, nil
}
