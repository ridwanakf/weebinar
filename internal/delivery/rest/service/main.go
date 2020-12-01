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

func GetUserSessionID(c echo.Context) (int64, error) {
	sess, err := session.Get("session", c)
	if err != nil {
		return -1, err
	}

	id, ok := sess.Values["id"]
	if !ok {
		return -1, entity.ErrIDNotFoundInSession
	}

	return id.(int64), nil
}