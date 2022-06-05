package server

import (
	"context"
	"io"
	"log"
	"net"
	"os"

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
	l := log.New(os.Stdout, "", 0)

	// Write telemetry data to a file.
	f, err := os.Create("traces.txt")
	if err != nil {
		l.Fatal(err)
	}
	defer f.Close()

	exp, err := newExporter(f)
	if err != nil {
		l.Fatal(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tp)
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
