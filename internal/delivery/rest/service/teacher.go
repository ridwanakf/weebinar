package service

import (
	"github.com/labstack/echo"
	"github.com/ridwanakf/weebinar/internal"
	"github.com/ridwanakf/weebinar/internal/app"
	"net/http"
)

type TeacherService struct {
	uc internal.TeacherUC
}

func NewTeacherService(app *app.WeebinarApp) *TeacherService {
	return &TeacherService{
		uc: app.UseCases.TeacherUC,
	}
}

func (s *TeacherService) IndexHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
