package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gmx-delta-neutral/gmx-neutral.query/internal/infrastructure"
	"github.com/gmx-delta-neutral/gmx-neutral.query/pkg/api"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

const ServiceName = "gmx-neutral.query-client"

type asset struct {
	Symbol     string
	Weight     *big.Int
	PoolAmount *big.Int
}

func main() {
	logger := infrastructure.NewLogger(ServiceName)
	interceptors := infrastructure.NewInterceptors(logger)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()), grpc.WithUnaryInterceptor(interceptors.TracingClientInterceptor(ServiceName)))
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()

	c := api.NewPositionServiceClient(conn)
	g := api.NewGlpServiceClient(conn)
	p := api.NewPriceServiceClient(conn)

	// create request
	req := api.GetTokenPositionRequest{
		TokenAddress: "0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664",
	}

	// call Greet service
	res, err := c.GetTokenPosition(context.Background(), &req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	fmt.Println(res.TokenPosition)

	glpReq := api.GetGlpAssetsRequest{}

	glpRes, err := g.GetGlpAssets(context.Background(), &glpReq)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	fmt.Println(glpRes)

	tokenPriceReq := api.GetTokenPriceRequest{
		Address: common.HexToAddress("0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab").Bytes(),
	}

	// call Greet service
	tokenPriceRes, err := p.GetTokenPrice(context.Background(), &tokenPriceReq)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	price := new(big.Int).SetBytes(tokenPriceRes.Price)
	fmt.Println(price)

	tokenPositionReq := api.GetTokenPositionRequest{
		TokenAddress: "0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664",
	}

	tokenPositionRes, err := c.GetTokenPosition(context.Background(), &tokenPositionReq)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	fmt.Println(new(big.Int).SetBytes(tokenPositionRes.TokenPosition.Amount).String())
	fmt.Println(new(big.Int).SetBytes(tokenPositionRes.TokenPosition.Pnl).String())
	fmt.Println(new(big.Int).SetBytes(tokenPositionRes.TokenPosition.Worth).String())
}
