package rest

import (
	"github.com/labstack/echo"
	"github.com/ridwanakf/weebinar/internal/app"
	"github.com/ridwanakf/weebinar/internal/delivery/rest/server"
	"github.com/ridwanakf/weebinar/internal/delivery/rest/service"
)

func initAPIHandler(eg *echo.Group, svc *service.Services) {
}

func Start(app *app.WeebinarApp) {
	srv := server.New()
	svc := service.GetServices(app)

	// API Handlers
	api := srv.Group("/api/v1")
	initAPIHandler(api, svc)

	server.Start(srv, &app.Cfg.Server)
}
