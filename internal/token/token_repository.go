package token

import (
	"context"
	"errors"
	"math/big"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/model"
)

type Repository interface {
	GetTokenPositions() ([]*model.TokenPosition, error)
	GetTokenPosition(address common.Address) (*model.TokenPosition, error)
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
	return &tokenRepository{}
}

var walletAddress = common.HexToAddress("0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664")

func (repository *tokenRepository) GetTokenPosition(ercAddress common.Address) (*model.TokenPosition, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
	)

	if err != nil {
		return nil, err
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
		return nil, err
	}

	positionModel := &TokenPositionModel{}

	err = attributevalue.UnmarshalMap(out.Item, positionModel)

	if err != nil {
		return nil, err
	}

	balance, success := new(big.Int).SetString(positionModel.Balance, 10)

	if !success {
		return nil, errors.New("Not successful")
	}

	position := &model.TokenPosition{
		TokenAddress:  common.HexToAddress(positionModel.TokenAddress),
		WalletAddress: common.HexToAddress(positionModel.TokenAddress),
		Symbol:        positionModel.Symbol,
		Balance:       *balance,
		Decimals:      positionModel.Decimals,
	}

	return position, nil
}

func (repository *tokenRepository) GetTokenPositions() ([]*model.TokenPosition, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
	)

	if err != nil {
		return nil, err
	}

	svc := dynamodb.NewFromConfig(cfg)

	out, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("TokenPositions"),
	})

	if err != nil {
		return nil, err
	}

	positionsModel := []*TokenPositionModel{}

	err = attributevalue.UnmarshalListOfMaps(out.Items, positionsModel)

	if err != nil {
		return nil, err
	}

	positions := []*model.TokenPosition{}

	for _, positionModel := range positionsModel {
		balance, success := new(big.Int).SetString(positionModel.Balance, 10)

		if !success {
			return nil, errors.New("Not successful")
		}

		position := &model.TokenPosition{
			TokenAddress:  common.HexToAddress(positionModel.TokenAddress),
			WalletAddress: common.HexToAddress(positionModel.TokenAddress),
			Symbol:        positionModel.Symbol,
			Balance:       *balance,
			Decimals:      positionModel.Decimals,
		}

		positions = append(positions, position)
	}

	return positions, nil
}
