package token

import (
	"github.com/ethereum/go-ethereum/common"
	cfg "github.com/gmx-delta-neutral/gmx-neutral.query/internal/config"
)

type Repository interface {
	// GetTokenPositions() ([]*model.TokenPosition, error)
	// GetTokenTransactions(address common.Address) ([]*model.Transaction, error)
}

type TransactionModel struct {
	TransactionId string
	TokenAddress  string
	WalletAddress string
	Amount        string
	PurchasePrice string
}

type tokenRepository struct {
	config *cfg.Config
}

func NewRepository(config *cfg.Config) Repository {
	return &tokenRepository{
		config: config,
	}
}

var walletAddress = common.HexToAddress("0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664")

/*
func (repository *tokenRepository) GetTokenTransactions(ercAddress common.Address) ([]*model.Transaction, error) {

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
		TableName:        aws.String("TokenPositions"),
		FilterExpression: aws.String("TokenAddress = :TokenAddress"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":TokenAddress": &types.AttributeValueMemberS{Value: ercAddress.String()},
		},
	})

	if err != nil {
		return nil, err
	}

	transactionModels := []*TransactionModel{}

	err = attributevalue.UnmarshalListOfMaps(out.Items, &transactionModels)

	if err != nil {
		return nil, err
	}

	transactions := []*model.Transaction{}

	for _, transactionModel := range transactionModels {
		amount, successAmount := new(big.Int).SetString(transactionModel.Amount, 10)
		purchasePrice, successPurchasePrice := new(big.Int).SetString(transactionModel.PurchasePrice, 10)

		if !successAmount || !successPurchasePrice {
			return nil, errors.New("not successful parsing bigint")
		}

		transactions = append(transactions, &model.Transaction{
			TransactionId: common.HexToHash(transactionModel.TransactionId),
			TokenAddress:  common.HexToAddress(transactionModel.TokenAddress),
			WalletAddress: common.HexToAddress(transactionModel.TokenAddress),
			Amount:        amount,
			PurchasePrice: purchasePrice,
		})
	}

	return transactions, nil
}
*/

// func (repository *tokenRepository) GetTokenPositions() ([]*model.TokenPosition, error) {
// 	cfg, err := config.LoadDefaultConfig(context.Background(),
// 		config.WithRegion("us-east-1"),
// 		config.WithEndpointResolver(aws.EndpointResolverFunc(
// 			func(service, region string) (aws.Endpoint, error) {
// 				return aws.Endpoint{URL: "http://localhost:8000"}, nil
// 			})),
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	svc := dynamodb.NewFromConfig(cfg)

// 	out, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
// 		TableName: aws.String("TokenPositions"),
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	positionsModel := []*TokenPositionModel{}

// 	err = attributevalue.UnmarshalListOfMaps(out.Items, &positionsModel)

// 	if err != nil {
// 		return nil, err
// 	}

// 	positions := []*model.TokenPosition{}

// 	for _, positionModel := range positionsModel {
// 		balance, success := new(big.Int).SetString(positionModel.Balance, 10)

// 		if !success {
// 			return nil, errors.New("Not successful")
// 		}

// 		position := &model.TokenPosition{
// 			TokenAddress:  common.HexToAddress(positionModel.TokenAddress),
// 			WalletAddress: common.HexToAddress(positionModel.TokenAddress),
// 			Symbol:        positionModel.Symbol,
// 			Balance:       *balance,
// 			Decimals:      positionModel.Decimals,
// 		}

// 		positions = append(positions, position)
// 	}

// 	return positions, nil
// }
