package rest

import (
	"github.com/labstack/echo"
	"github.com/ridwanakf/weebinar/internal/app"
	"github.com/ridwanakf/weebinar/internal/delivery/rest/server"
	"github.com/ridwanakf/weebinar/internal/delivery/rest/service"
)

func initCommonHandlers(eg *echo.Group, svc *service.Services) {
	eg.GET("", svc.CommonService.IndexHandler)
}

func initTeacherHandlers(eg *echo.Group, svc *service.Services) {
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
	initTeacherHandlers(teacherGroup, svc)

	studentGroup := srv.Group("/student")
	initStudentHandlers(studentGroup, svc)

	server.Start(srv, &app.Cfg.Server)
}
