package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/db"
	walletService "gitlab.rinznetwork.com/gocryptowallet/go-template/internal/domains/wallets/service"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/group"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/middlewares"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/interceptors"
	kafkaClient "gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/kafka"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/postgres"
)

type server struct {
	log      logger.Logger
	cfg      *config.Config
	v        *validator.Validate
	server   *http.Server
	ws       *walletService.WalletService
	im       interceptors.InterceptorManager
	writerDB *db.Store
	readerDB *db.Store
	mw       middlewares.MiddlewareManager
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg, v: validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.im = interceptors.NewInterceptorManager(s.log)

	readerPgxConn, err := postgres.NewPgxConn(s.cfg.READER_POSTGRES_URL)
	if err != nil {
		return errors.Wrap(err, "postgresql.NewPgxConn")
	}
	s.readerDB = db.NewStore(readerPgxConn)

	s.log.Infof("reader postgres connected: %v", readerPgxConn.Stat().TotalConns())
	defer readerPgxConn.Close()

	writerPgxConn, err := postgres.NewPgxConn(s.cfg.WRITER_POSTGRES_URL)
	if err != nil {
		return errors.Wrap(err, "postgresql.NewPgxConn")
	}
	s.writerDB = db.NewStore(writerPgxConn)

	s.log.Infof("writer postgres connected: %v", writerPgxConn.Stat().TotalConns())
	defer writerPgxConn.Close()

	kafkaProducer := kafkaClient.NewProducer(
		s.log,
		s.cfg.KAFKA_BROKER,
		s.cfg.AWS_REGION,
		s.cfg.AWS_ACCESS_KEY_ID,
		s.cfg.AWS_SECRET_ACCESS_KEY,
	)
	defer kafkaProducer.Close() // nolint: errcheck

	s.server = &http.Server{
		Addr: ":" + s.cfg.PORT,
	}

	router := mux.NewRouter().PathPrefix(s.cfg.PREFIX_PATH).Subrouter()

	group.InitGroup(router, s.log, s.mw, s.cfg, s.v, ctx, kafkaProducer, s.writerDB, s.readerDB)

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
