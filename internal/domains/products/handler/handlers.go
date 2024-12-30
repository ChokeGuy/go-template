package handler

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"

	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/products/service"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/middlewares"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type productsHandlers struct {
	router *mux.Router
	log    logger.Logger
	mw     middlewares.MiddlewareManager
	cfg    *config.Config
	ps     *service.ProductService
	v      *validator.Validate
	ctx    context.Context
}

func NewProductsHandlers(
	router *mux.Router,
	log logger.Logger,
	mw middlewares.MiddlewareManager,
	cfg *config.Config,
	ps *service.ProductService,
	v *validator.Validate,
	ctx context.Context,
) *productsHandlers {
	return &productsHandlers{router: router, log: log, mw: mw, cfg: cfg, ps: ps, v: v, ctx: ctx}
}
