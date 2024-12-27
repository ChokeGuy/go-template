package server

import (
	"time"

	"gitlab.rinznetwork.com/gocryptowallet/go-template/docs"
)

const (
	maxHeaderBytes = 1 << 20
	stackSize      = 1 << 10 // 1 KB
	bodyLimit      = "2M"
	readTimeout    = 15 * time.Second
	writeTimeout   = 15 * time.Second
	gzipLevel      = 5
)

func (s *server) runHttpServer() error {
	// s.mapRoutes()

	s.server.ReadTimeout = readTimeout
	s.server.WriteTimeout = writeTimeout
	s.server.MaxHeaderBytes = maxHeaderBytes

	return s.server.ListenAndServe()
}

func (s *server) mapRoutes() {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "API Gateway"
	docs.SwaggerInfo.Description = "API Gateway CQRS microservices."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"

	// s.server.Handler.GET("/swagger/*", echoSwagger.WrapHandler)

	// s.echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
	// 	StackSize:         stackSize,
	// 	DisablePrintStack: true,
	// 	DisableStackAll:   true,
	// }))
	// s.echo.Use(middleware.RequestID())
	// s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
	// 	Level: gzipLevel,
	// 	Skipper: func(c echo.Context) bool {
	// 		return strings.Contains(c.Request().URL.Path, "swagger")
	// 	},
	// }))
	// s.echo.Use(middleware.BodyLimit(bodyLimit))
}
