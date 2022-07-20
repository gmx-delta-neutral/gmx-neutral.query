package transaction

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	cfg "github.com/gmx-delta-neutral/gmx-neutral.query/internal/config"
	glpmanager "github.com/gmx-delta-neutral/gmx-neutral.query/internal/contracts/glp_manager"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/model"
)

type Service interface {
	GetGlpTransactions() (*[]*model.Transaction, error)
}

type TransactionService struct {
	ethClient *ethclient.Client
	config    *cfg.Config
}

func NewTransactionService(config *cfg.Config, ethClient *ethclient.Client) Service {
	return &TransactionService{
		config:    config,
		ethClient: ethClient,
	}
}

type snowtraceTransactionResult struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Result  []snowtraceTransaction `json:"result"`
}

type snowtraceTransaction struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxreceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
	MethodID          string `json:"methodId"`
	FunctionName      string `json:"functionName"`
}

func (s *TransactionService) GetGlpTransactions() (*[]*model.Transaction, error) {
	walletAddress := common.HexToAddress(s.config.Addresses.Account)
	req, err := http.NewRequest("GET", "https://api.snowtrace.io/api", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("module", "account")
	q.Add("action", "txlist")
	q.Add("address", walletAddress.String())
	q.Add("apiKey", s.config.SnowtraceApiKey)
	q.Add("sort", "desc")

	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	snowtraceResult := &snowtraceTransactionResult{}

	json.NewDecoder(res.Body).Decode(snowtraceResult)

	if err != nil {
		return nil, err
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(glpmanager.GlpmanagerABI)))

	if err != nil {
		log.Fatal(err)
	}

	transactions := []*model.Transaction{}

	for _, snowTraceTransaction := range snowtraceResult.Result {

		if snowTraceTransaction.To != s.config.Addresses.RewardRouter {
			continue
		}

		transactionHash := common.HexToHash(snowTraceTransaction.Hash)
		receipt, err := s.ethClient.TransactionReceipt(context.Background(), transactionHash)

		if err != nil {
			return nil, err
		}

		for _, log := range receipt.Logs {
			if !s.isAddLiquidityLog(log) {
				continue
			}

			data, err := contractAbi.Unpack("AddLiquidity", log.Data)

			if err != nil {
				return nil, err
			}

			usdAmount := data[2]
			mintAmount := data[6]

			transaction := &model.Transaction{
				TransactionId: transactionHash,
				TokenAddress:  common.HexToAddress(s.config.Addresses.Glp),
				WalletAddress: walletAddress,
				Symbol:        "GLP",
				Amount:        mintAmount.(*big.Int),
				Decimals:      int(s.config.Decimals.Glp),
				CostBasis:     usdAmount.(*big.Int),
			}

			transactions = append(transactions, transaction)
		}

		return &transactions, nil

	}

	return &[]*model.Transaction{}, nil
}

func (s *TransactionService) isAddLiquidityLog(log *types.Log) bool {
	for _, topic := range log.Topics {
		if topic.String() == s.config.Hashes.AddLiquidity {
			return true
		}
	}

	return false
}
