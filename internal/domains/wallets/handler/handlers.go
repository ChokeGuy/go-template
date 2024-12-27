package handler

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"

	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/service"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/middlewares"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type walletsHandlers struct {
	router *mux.Router
	log    logger.Logger
	mw     middlewares.MiddlewareManager
	cfg    *config.Config
	ws     *service.WalletService
	v      *validator.Validate
	ctx    context.Context
}

func NewWalletsHandlers(
	router *mux.Router,
	log logger.Logger,
	mw middlewares.MiddlewareManager,
	cfg *config.Config,
	ws *service.WalletService,
	v *validator.Validate,
	ctx context.Context,
) *walletsHandlers {

	return &walletsHandlers{router: router, log: log, mw: mw, cfg: cfg, ws: ws, v: v, ctx: ctx}
}
