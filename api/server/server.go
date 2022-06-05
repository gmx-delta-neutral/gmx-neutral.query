package server

import (
	"log"
	"net"

	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/api/generated"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/infrastructure"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/position"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	logger := infrastructure.NewLogger()
	interceptors := infrastructure.NewInterceptors(logger)

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(otelgrpc.UnaryServerInterceptor(), interceptors.LoggingInterceptor()))
	positionServer := NewPositionServer(position.NewService(position.NewRepository()))
	generated.RegisterPositionServiceServer(s, positionServer)

	log.Println("Listening on port 50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

}
