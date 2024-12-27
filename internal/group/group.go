package group

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/db"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/middlewares"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

func InitGroup(
	router *mux.Router,
	log logger.Logger,
	mw middlewares.MiddlewareManager,
	cfg *config.Config,
	v *validator.Validate,
	ctx context.Context,
	// kafkaProducer kafkaClient.Producer,
	writerDB *db.Store,
	readerDB *db.Store,
) {
	walletGroup(router, log, mw, cfg, v, ctx, writerDB, readerDB)
}
