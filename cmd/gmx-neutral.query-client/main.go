package main

import (
	"context"
	"fmt"
	"log"

	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/api/generated"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/infrastructure"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

const ServiceName = "gmx-neutral.query-client"

func main() {
	logger := infrastructure.NewLogger(ServiceName)
	interceptors := infrastructure.NewInterceptors(logger)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()), grpc.WithUnaryInterceptor(interceptors.TracingClientInterceptor(ServiceName)))
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()

	c := generated.NewPositionServiceClient(conn)

	// create request
	req := generated.GetPositionsRequest{}

	// call Greet service
	res, err := c.GetPositions(context.Background(), &req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	fmt.Println(res.Positions)
}
