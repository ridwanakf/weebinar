package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ridwanakf/weebinar/internal"
	"github.com/ridwanakf/weebinar/internal/app"
)

type TeacherService struct {
	uc internal.TeacherUC
}

func NewTeacherService(app *app.WeebinarApp) *TeacherService {
	return &TeacherService{
		uc: app.UseCases.TeacherUC,
	}
}

func (s *TeacherService) HomeTeacherPageHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home_teacher", nil)
}

func (s *TeacherService) ProfileTeacherPageHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "settings", nil)
}

func (s *TeacherService) CreateWebinarPageHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "create", nil)
}

func (s *TeacherService) CreateWebinarPostHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}