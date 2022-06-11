package server

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/api/generated"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/infrastructure"
	"github.com/RafGDev/gmx-delta-neutral/gmx-neutral.query/internal/position"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
)

const ServiceName = "gmx-neutral.query"

func Resource() *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("stdout-example"),
		semconv.ServiceVersionKey.String("0.0.1"),
	)
}

func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(ServiceName),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "dev"),
		),
	)
	return r
}

func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

func StartServer() {
	ctx := context.Background()
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		panic(err)
	}

	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("gmx-neutral.query"),
		semconv.ServiceVersionKey.String("1.0.0"),
		semconv.DeploymentEnvironmentKey.String("production"),
	)

	provider := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource),
	)

	defer func() {
		if err := provider.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	}()

	otel.SetTracerProvider(provider)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	logger := infrastructure.NewLogger(ServiceName)
	interceptors := infrastructure.NewInterceptors(logger)

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(otelgrpc.UnaryServerInterceptor(), interceptors.LoggingServerInterceptor()))
	positionServer := NewPositionServer(position.NewService(position.NewRepository()))
	generated.RegisterPositionServiceServer(s, positionServer)

	log.Println("Listening on port 50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

}
