package infrastructure

import (
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otellogrus"
)

func NewLogger() *log.Entry {
	log.SetFormatter(&log.JSONFormatter{})

	logger := log.WithFields(log.Fields{
		"service": "gmx-neutral.query",
	})

	log.AddHook(otellogrus.NewHook(otellogrus.WithLevels(
		log.AllLevels...,
	)))

	return logger
}
