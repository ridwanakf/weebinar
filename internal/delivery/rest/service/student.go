package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ridwanakf/weebinar/internal"
	"github.com/ridwanakf/weebinar/internal/app"
)

type StudentService struct {
	uc internal.StudentUC
}

func NewStudentService(app *app.WeebinarApp) *StudentService {
	return &StudentService{
		uc: app.UseCases.StudentUC,
	}
}

func (s *StudentService) IndexHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
