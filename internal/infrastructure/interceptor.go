package infrastructure

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type interceptors struct {
	logger *log.Entry
}

func NewInterceptors(logger *log.Entry) *interceptors {
	return &interceptors{logger}
}

func (i *interceptors) LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()

		i.logger.WithContext(ctx).
			WithFields(log.Fields{
				"method": info.FullMethod,
			}).Info("Request started")

		h, err := handler(ctx, req)

		i.logger.WithContext(ctx).
			WithFields(log.Fields{
				"method":   info.FullMethod,
				"duration": time.Since(start),
			}).Info("Request finished")

		return h, err
	}
}
