package service

import (
	"github.com/labstack/echo"
	"github.com/ridwanakf/weebinar/internal"
	"github.com/ridwanakf/weebinar/internal/app"
	"net/http"
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
