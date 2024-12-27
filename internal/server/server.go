package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/segmentio/kafka-go"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	productHandler "gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/products/handler"
	productService "gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/products/service"
	walletHandler "gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/handler"
	walletService "gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/service"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/middlewares"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/interceptors"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

type server struct {
	log       logger.Logger
	cfg       *config.Config
	v         *validator.Validate
	server    *http.Server
	kafkaConn *kafka.Conn
	ps        *productService.ProductService
	ws        *walletService.WalletService
	im        interceptors.InterceptorManager
	pgConn    *pgxpool.Pool
	mw        middlewares.MiddlewareManager
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg, v: validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.im = interceptors.NewInterceptorManager(s.log)

	// pgxConn, err := postgres.NewPgxConn(s.cfg.POSTGRES_URL)
	// if err != nil {
	// 	return errors.Wrap(err, "postgresql.NewPgxConn")
	// }
	// s.pgConn = pgxConn
	// s.log.Infof("postgres connected: %v", pgxConn.Stat().TotalConns())
	// defer pgxConn.Close()

	// kafkaProducer := kafkaClient.NewProducer(s.log, s.cfg.KAFKA_BROKER)
	// defer kafkaProducer.Close() // nolint: errcheck

	// if err := s.connectKafkaBrokers(ctx); err != nil {
	// 	return errors.Wrap(err, "s.connectKafkaBrokers")
	// }
	// defer s.kafkaConn.Close() // nolint: errcheck

	// if s.cfg.INIT_TOPICS {
	// 	s.initKafkaTopics(ctx)
	// }

	s.server = &http.Server{
		Addr: ":8080",
	}

	// s.ps = productService.NewProductService(s.log, s.cfg, kafkaProducer)
	s.ws = walletService.NewWalletService(s.log, s.cfg)

	router := mux.NewRouter().PathPrefix(s.cfg.PREFIX_PATH).Subrouter()

	productHandler.NewProductsHandlers(router, s.log, s.mw, s.cfg, s.ps, s.v, ctx).MapRoutes()
	walletHandler.NewWalletsHandlers(router, s.log, s.mw, s.cfg, s.ws, s.v, ctx).MapRoutes()

	s.server.Handler = router

	go func() {
		if err := s.runHttpServer(); err != nil {
			s.log.Errorf(" s.runHttpServer: %v", err)
			cancel()
		}
	}()
	s.log.Infof("API Gateway is listening on PORT: %s", s.cfg.PORT)

	<-ctx.Done()
	if err := s.server.Shutdown(ctx); err != nil {
		s.log.WarnMsg("Server.Shutdown", err)
	}

	return nil
}
