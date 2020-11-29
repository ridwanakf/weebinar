package service

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ridwanakf/weebinar/internal/app"
)

type CommonService struct {
}

func NewCommonService(app *app.WeebinarApp) *CommonService {
	return &CommonService{
	}
}

func (s *CommonService) IndexHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
