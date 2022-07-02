package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/api/generated"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/infrastructure"
	"github.com/ethereum/go-ethereum/common"
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

	c := generated.NewPositionServiceClient(conn)
	g := generated.NewGlpServiceClient(conn)
	p := generated.NewPriceServiceClient(conn)

	// create request
	req := generated.GetTokenPositionRequest{
		TokenAddress: "0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664",
	}

	// call Greet service
	res, err := c.GetTokenPosition(context.Background(), &req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	fmt.Println(res.TokenPosition)

	glpReq := generated.GetGlpAssetsRequest{}

	glpRes, err := g.GetGlpAssets(context.Background(), &glpReq)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	fmt.Println(glpRes)

	tokenPriceReq := generated.GetTokenPriceRequest{
		Address: common.HexToAddress("0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab").Bytes(),
	}

	// call Greet service
	tokenPriceRes, err := p.GetTokenPrice(context.Background(), &tokenPriceReq)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	price := new(big.Int).SetBytes(tokenPriceRes.Price)
	fmt.Println(price)
}
