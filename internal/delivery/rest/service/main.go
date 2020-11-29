package service

import "github.com/ridwanakf/weebinar/internal/app"

type Services struct {
	*CommonService
	*TeacherService
	*StudentService
}

func GetServices(app *app.WeebinarApp) *Services {
	return &Services{
		CommonService : NewCommonService(app),
		TeacherService: NewTeacherService(app),
		StudentService: NewStudentService(app),
	}
}
