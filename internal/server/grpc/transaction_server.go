package grpc_server

import (
	"context"

	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/transaction"
	"github.com/gmx-delta-neutral/gmx-neutral.query/pkg/api"
)

func NewTransactionServer(transactionService transaction.Service) *TransactionServer {
	return &TransactionServer{
		transactionService: transactionService,
	}
}

type TransactionServer struct {
	transactionService transaction.Service
}

func (p *TransactionServer) GetTransactions(ctx context.Context, request *api.GetTransactionsRequest) (*api.GetTransactionsResponse, error) {
	transactions, err := p.transactionService.GetTransactions()

	if err != nil {
		return nil, err
	}
	transactionsResponse := &api.GetTransactionsResponse{
		Transactions: TransactionDtos(transactions),
	}

	return transactionsResponse, nil
}
