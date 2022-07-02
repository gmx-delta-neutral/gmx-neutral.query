package price

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Service interface {
	GetPrice(address common.Address) (*big.Int, error)
}

type priceService struct {
	priceRepository Repository
}

type PriceFeedMapping struct {
	Address          common.Address
	PriceFeedAddress common.Address
}

var priceFeedMappings = []PriceFeedMapping{
	{
		Address:          common.HexToAddress("0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7"),
		PriceFeedAddress: common.HexToAddress("0x0A77230d17318075983913bC2145DB16C7366156"),
	},
	{
		Address:          common.HexToAddress("0x50b7545627a5162f82a992c33b87adc75187b218"),
		PriceFeedAddress: common.HexToAddress("0x2779D32d5166BAaa2B2b658333bA7e6Ec0C65743"),
	},
	{
		Address:          common.HexToAddress("0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab"),
		PriceFeedAddress: common.HexToAddress("0x976B3D034E162d8bD72D6b9C989d545b839003b0"),
	},
	{
		Address:          common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e"),
		PriceFeedAddress: common.HexToAddress("0xF096872672F44d6EBA71458D74fe67F9a77a23B9"),
	},
	{
		Address:          common.HexToAddress("0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664"),
		PriceFeedAddress: common.HexToAddress("0xF096872672F44d6EBA71458D74fe67F9a77a23B9"),
	},
}

func NewPriceService(priceRepository Repository) Service {
	return &priceService{
		priceRepository: priceRepository,
	}
}

func (s *priceService) GetPrice(address common.Address) (*big.Int, error) {
	priceFeedAddress, err := getFeedAddressFromTokenAddress(address)

	if err != nil {
		return nil, err
	}

	return s.priceRepository.GetPrice(priceFeedAddress)
}

func getFeedAddressFromTokenAddress(tokenAddress common.Address) (common.Address, error) {
	for _, mapping := range priceFeedMappings {
		if mapping.Address.String() == tokenAddress.String() {
			return mapping.PriceFeedAddress, nil
		}
	}

	return common.Address{}, errors.New("Error")
}
