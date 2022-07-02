package token

import (
	"context"
	"errors"
	"math/big"

	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ethereum/go-ethereum/common"
)

type Repository interface {
	GetTokenPosition(address common.Address) (model.TokenPosition, error)
}

type TokenPositionModel struct {
	TokenAddress  string
	WalletAddress string
	Symbol        string
	Balance       string
	Decimals      int32
}

type tokenRepository struct{}

func NewRepository() Repository {
	return tokenRepository{}
}

var walletAddress = common.HexToAddress("0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664")

func (repository tokenRepository) GetTokenPosition(ercAddress common.Address) (model.TokenPosition, error) {
	position := model.TokenPosition{}

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
	)

	if err != nil {
		return position, err
	}

	svc := dynamodb.NewFromConfig(cfg)

	out, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("TokenPositions"),
		Key: map[string]types.AttributeValue{
			"TokenAddress":  &types.AttributeValueMemberS{Value: ercAddress.String()},
			"WalletAddress": &types.AttributeValueMemberS{Value: walletAddress.String()},
		},
	})

	if err != nil {
		return position, err
	}

	positionModel := TokenPositionModel{}

	err = attributevalue.UnmarshalMap(out.Item, &positionModel)

	if err != nil {
		return position, err
	}

	balance, success := new(big.Int).SetString(positionModel.Balance, 10)

	if !success {
		return position, errors.New("Not successful")
	}

	position = model.TokenPosition{
		TokenAddress:  common.HexToAddress(positionModel.TokenAddress),
		WalletAddress: common.HexToAddress(positionModel.TokenAddress),
		Symbol:        positionModel.Symbol,
		Balance:       *balance,
		Decimals:      positionModel.Decimals,
	}

	return position, nil
}
