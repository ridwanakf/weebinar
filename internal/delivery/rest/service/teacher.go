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

func (s *TeacherService) HomeTeacherHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}

func (s *TeacherService) ProfileTeacherHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "settings", nil)
}

func (s *TeacherService) CreateWebinarHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "create", nil)
}

func (s *TeacherService) PostCreateWebinarHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}