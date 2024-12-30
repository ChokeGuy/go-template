package main

import (
	"flag"
	"log"

	"gitlab.rinznetwork.com/gocryptowallet/go-template/config"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/internal/server"
	"gitlab.rinznetwork.com/gocryptowallet/go-template/pkg/logger"
)

// @contact.name Alexander Bryksin
// @contact.url https://github.com/AleksK1NG
// @contact.email alexander.bryksin@yandex.ru
func main() {
	flag.Parse()

	cfg, err := config.InitConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger()
	appLogger.InitLogger()
	appLogger.WithName("ApiGateway")

	s := server.NewServer(appLogger, cfg)
	appLogger.Fatal(s.Run())
}
