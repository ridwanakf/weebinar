package service

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ridwanakf/weebinar/internal/app"
	"github.com/ridwanakf/weebinar/internal/entity"
)

type Services struct {
	*CommonService
	*TeacherService
	*StudentService
}

func GetServices(app *app.WeebinarApp) *Services {
	return &Services{
		CommonService:  NewCommonService(app),
		TeacherService: NewTeacherService(app),
		StudentService: NewStudentService(app),
	}
}

func GetUserSessionID(c echo.Context) (string, error) {
	sess, err := session.Get("session", c)
	if err != nil {
		return "", err
	}

	id, ok := sess.Values["id"]
	if !ok {
		return "", entity.ErrIDNotFoundInSession
	}

	return id.(string), nil
}