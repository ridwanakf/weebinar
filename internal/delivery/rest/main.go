package rest

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ridwanakf/weebinar/internal/app"
	"github.com/ridwanakf/weebinar/internal/delivery/rest/middleware"
	"github.com/ridwanakf/weebinar/internal/delivery/rest/server"
	"github.com/ridwanakf/weebinar/internal/delivery/rest/service"
	"net/http"
)

func initCommonHandlers(eg *echo.Group, svc *service.Services) {
	eg.GET("/", svc.CommonService.IndexHandler)
	eg.POST("/teacher/login", svc.CommonService.TeacherSignInHandler)
	eg.GET("/teacher/login/callback", svc.CommonService.TeacherSignInCallbackHandler)
	eg.POST("/student/login", svc.CommonService.StudentSignInHandler)
	eg.GET("/student/login/callback", svc.CommonService.StudentSignInCallbackHandler)

	// TODO: hapus kalo udah gk dipake buat test session lagi
	eg.GET("/sess", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["id"] = int64(12)
		sess.Values["role"] = "teacher"
		sess.Save(c.Request(), c.Response())
		return c.NoContent(http.StatusOK)
	})
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
	commonGroup := srv.Group("")
	initCommonHandlers(commonGroup, svc)

	teacherGroup := srv.Group("/teacher")
	teacherGroup.Use(middleware.CheckSession())
	initTeacherHandlers(teacherGroup, svc)

	studentGroup := srv.Group("/student")
	studentGroup.Use(middleware.CheckSession())
	initStudentHandlers(studentGroup, svc)

	server.Start(srv, &app.Cfg.Server)
}
