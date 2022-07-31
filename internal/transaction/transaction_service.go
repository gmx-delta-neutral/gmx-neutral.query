package transaction

import (
	"context"
	"encoding/json"
	"fmt"
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
	balancervault "github.com/gmx-delta-neutral/gmx-neutral.query/internal/contracts/balancer_vault"
	glpmanager "github.com/gmx-delta-neutral/gmx-neutral.query/internal/contracts/glp_manager"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/model"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/price"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/util"
)

type Service interface {
	GetTransactions() ([]*model.Transaction, error)
	GetGlpTransactions() ([]*model.Transaction, error)
	GetTracerTransactions() ([]*model.Transaction, error)
}

type TransactionService struct {
	ethClient       *ethclient.Client
	arbitrumClient  *ethclient.Client
	config          *cfg.Config
	priceRepository price.Repository
}

func NewTransactionService(config *cfg.Config, ethClient *ethclient.Client, arbitrumClient *ethclient.Client, priceRepository price.Repository) Service {
	return &TransactionService{
		config:          config,
		ethClient:       ethClient,
		arbitrumClient:  arbitrumClient,
		priceRepository: priceRepository,
	}
}

type blockExplorerTransactionResult struct {
	Status  string                     `json:"status"`
	Message string                     `json:"message"`
	Result  []blockExplorerTransaction `json:"result"`
}

type blockExplorerTransaction struct {
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

type tracerPendingCommit struct {
	Amount           string `json:"amount"`
	CommitType       int    `json:"commitType"`
	Type             string `json:"type"`
	TxnHash          string `json:"txnHash"`
	Timestamp        int    `json:"timestamp"`
	UserAddress      string `json:"userAddress"`
	CommitID         string `json:"commitID"`
	PoolAddress      string `json:"poolAddress"`
	UpdateIntervalID int    `json:"updateIntervalId"`
}

type TracerTradeHistoryResult struct {
	Page         int                 `json:"page"`
	PageSize     int                 `json:"pageSize"`
	TotalRecords int                 `json:"totalRecords"`
	Rows         []TracerTransaction `json:"rows"`
}

type TracerTransaction struct {
	Date                 int    `json:"date"`
	Pool                 string `json:"pool"`
	UserAddress          string `json:"userAddress"`
	Type                 string `json:"type"`
	UpdateIntervalID     string `json:"updateIntervalId"`
	PayForClaim          bool   `json:"payForClaim"`
	FromAggregateBalance bool   `json:"fromAggregateBalance"`
	TransactionHashIn    string `json:"transactionHashIn"`
	TransactionHashOut   string `json:"transactionHashOut"`
	TokenDecimals        int    `json:"tokenDecimals"`
	TokenInAddress       string `json:"tokenInAddress"`
	TokenInSymbol        string `json:"tokenInSymbol"`
	TokenInName          string `json:"tokenInName"`
	TokenInAmount        string `json:"tokenInAmount"`
	PriceIn              string `json:"priceIn"`
	PriceOut             string `json:"priceOut"`
	PriceTokenAddress    string `json:"priceTokenAddress"`
	PriceTokenName       string `json:"priceTokenName"`
	PriceTokenSymbol     string `json:"priceTokenSymbol"`
	Fee                  string `json:"fee"`
	MintingFee           string `json:"mintingFee"`
	BurningFee           string `json:"burningFee"`
	TokenOutAddress      string `json:"tokenOutAddress"`
	TokenOutSymbol       string `json:"tokenOutSymbol"`
	TokenOutName         string `json:"tokenOutName"`
	TokenOutAmount       string `json:"tokenOutAmount"`
}

func (s *TransactionService) GetTransactions() ([]*model.Transaction, error) {
	glpTransactions, err := s.GetGlpTransactions()

	if err != nil {
		return nil, err
	}

	tracerTransactions, err := s.GetTracerTransactions()

	if err != nil {
		return nil, err
	}

	transactions := append(glpTransactions, tracerTransactions...)
	return transactions, err
}

func (s *TransactionService) GetGlpTransactions() ([]*model.Transaction, error) {
	walletAddress := common.HexToAddress(s.config.Addresses.Account)
	req, err := http.NewRequest("GET", s.config.ThirdPartyApi.SnowtraceApiUrl, nil)
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

	snowtraceResult := &blockExplorerTransactionResult{}

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

		return transactions, nil

	}

	return []*model.Transaction{}, nil
}

func (s *TransactionService) GetBalancerTransactions() ([]*model.Transaction, error) {
	walletAddress := common.HexToAddress(s.config.Addresses.Account)
	req, err := http.NewRequest("GET", s.config.ThirdPartyApi.ArbiscanApiUrl, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("module", "account")
	q.Add("action", "txlist")
	q.Add("address", walletAddress.String())
	q.Add("apiKey", s.config.ArbiscanApiKey)
	q.Add("sort", "desc")

	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	snowtraceResult := &blockExplorerTransactionResult{}

	json.NewDecoder(res.Body).Decode(snowtraceResult)

	if err != nil {
		return nil, err
	}

	contractAbi, err := abi.JSON(strings.NewReader(balancervault.BalancervaultABI))

	if err != nil {
		log.Fatal(err)
	}

	transactions := []*model.Transaction{}

	for _, snowTraceTransaction := range snowtraceResult.Result {

		if strings.ToLower(snowTraceTransaction.To) != strings.ToLower(s.config.Addresses.BalancerVault) {
			continue
		}

		transactionHash := common.HexToHash(snowTraceTransaction.Hash)
		receipt, err := s.arbitrumClient.TransactionReceipt(context.Background(), transactionHash)

		if err != nil {
			return nil, err
		}

		for _, log := range receipt.Logs {
			if !s.isSwapLog(log) {
				continue
			}

			data, err := contractAbi.Unpack("Swap", log.Data)

			if err != nil {
				return nil, err
			}

			usdcIn := data[0]
			shortBtcOut := data[1]

			transaction := &model.Transaction{
				TransactionId: transactionHash,
				TokenAddress:  common.HexToAddress(s.config.Addresses.Glp),
				WalletAddress: walletAddress,
				Symbol:        "BtcShort3x",
				Amount:        shortBtcOut.(*big.Int),
				Decimals:      int(s.config.Decimals.Glp),
				CostBasis:     usdcIn.(*big.Int),
			}

			transactions = append(transactions, transaction)
		}
	}

	return transactions, nil
}

func (s *TransactionService) GetTracerTransactions() ([]*model.Transaction, error) {
	claimedTransactions, err := s.getTracerClaimedTransactions()
	if err != nil {
		return nil, err
	}
	pendingTransactions, err := s.getTracerPendingTransactions()
	if err != nil {
		return nil, err
	}

	return append(claimedTransactions, pendingTransactions...), nil
}

func (s *TransactionService) getTracerClaimedTransactions() ([]*model.Transaction, error) {
	walletAddress := common.HexToAddress(s.config.Addresses.Account)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/poolsv2/tradeHistory", s.config.ThirdPartyApi.TracerApiUrl), nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("userAddress", walletAddress.String())
	q.Add("pageSize", "1000")
	q.Add("sortDirection", "DESC")
	q.Add("network", "42161")
	q.Add("poolAddress", s.config.Addresses.LeveragedPool)

	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)

	fmt.Println(res.Body)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	tradeHistoryResult := &TracerTradeHistoryResult{}

	json.NewDecoder(res.Body).Decode(tradeHistoryResult)

	transactions := []*model.Transaction{}

	for _, tracerTransaction := range tradeHistoryResult.Rows {
		shortBtcAmount, _ := new(big.Int).SetString(tracerTransaction.TokenOutAmount, 10)
		usdcAmount, _ := new(big.Int).SetString(tracerTransaction.TokenInAmount, 10)

		transaction := &model.Transaction{
			TransactionId: common.HexToHash(tracerTransaction.TransactionHashOut),
			TokenAddress:  common.HexToAddress(s.config.Addresses.ShortBtcToken),
			WalletAddress: walletAddress,
			Symbol:        "BtcShort3x",
			Amount:        shortBtcAmount,
			Decimals:      int(tracerTransaction.TokenDecimals),
			CostBasis:     usdcAmount,
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (s *TransactionService) getTracerPendingTransactions() ([]*model.Transaction, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/poolsv2/pendingCommits", s.config.ThirdPartyApi.TracerApiUrl), nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("userAddress", s.config.Addresses.Account)
	q.Add("network", "42161")
	q.Add("poolAddress", s.config.Addresses.LeveragedPool)

	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	pendingCommits := []*tracerPendingCommit{}
	json.NewDecoder(res.Body).Decode(&pendingCommits)

	transactions := []*model.Transaction{}

	nextShortPrice, err := s.priceRepository.GetTracerNextShortBtcPrice()

	if err != nil {
		return nil, err
	}

	for _, pendingCommit := range pendingCommits {
		usdcAmount, _ := new(big.Int).SetString(pendingCommit.Amount, 10)
		tokensAmount := new(big.Int).Div(util.ExpandDecimals(usdcAmount, 6), nextShortPrice)
		transaction := &model.Transaction{
			TransactionId: common.HexToHash(pendingCommit.TxnHash),
			TokenAddress:  common.HexToAddress(s.config.Addresses.ShortBtcToken),
			WalletAddress: common.HexToAddress(s.config.Addresses.Account),
			Symbol:        "BtcShort3x",
			Amount:        tokensAmount,
			Decimals:      6,
			CostBasis:     usdcAmount,
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (s *TransactionService) isAddLiquidityLog(log *types.Log) bool {
	for _, topic := range log.Topics {
		if topic.String() == s.config.Hashes.AddLiquidity {
			return true
		}
	}

	return false
}

func (s *TransactionService) isSwapLog(log *types.Log) bool {
	for _, topic := range log.Topics {
		if topic.String() == s.config.Hashes.Swap {
			return true
		}
	}

	return false
}

func (s *TransactionService) isMintLong(log *types.Log) bool {
	for _, topic := range log.Topics {
		fmt.Printf("topic: %s\n", topic.String())
		if topic.String() == s.config.Hashes.CreateCommit {
			return true
		}
	}

	return false
}
