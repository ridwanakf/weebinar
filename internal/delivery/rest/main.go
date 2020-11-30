package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/ridwanakf/weebinar/internal/app"
	"github.com/ridwanakf/weebinar/internal/delivery/rest/middleware"
	"github.com/ridwanakf/weebinar/internal/delivery/rest/server"
	"github.com/ridwanakf/weebinar/internal/delivery/rest/service"
)

func initCommonHandlers(eg *echo.Group, svc *service.Services) {
	eg.GET("", svc.CommonService.IndexHandler)
}

func initTeacherHandlers(eg *echo.Group, svc *service.Services) {
	eg.GET("/home", svc.HomeTeacherPageHandler)
	eg.GET("/profile", svc.ProfileTeacherPageHandler)
	eg.GET("/create", svc.CreateWebinarPageHandler)
	eg.POST("/create", svc.CreateWebinarPostHandler)
}

func initStudentHandlers(eg *echo.Group, svc *service.Services) {
}

func Start(app *app.WeebinarApp) {
	srv := server.New()
	svc := service.GetServices(app)

	// Handlers Groups
	commonGroup := srv.Group("/")
	initCommonHandlers(commonGroup, svc)

	teacherGroup := srv.Group("/teacher")
	teacherGroup.Use(middleware.CheckSession())
	initTeacherHandlers(teacherGroup, svc)

	studentGroup := srv.Group("/student")
	studentGroup.Use(middleware.CheckSession())
	initStudentHandlers(studentGroup, svc)

	server.Start(srv, &app.Cfg.Server)
}
