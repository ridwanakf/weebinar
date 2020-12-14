package service

import (
	"github.com/ridwanakf/weebinar/internal/entity"
	"log"
	"net/http"
	"strconv"

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

func (s *StudentService) HomeStudentPageHandler(c echo.Context) error {
	id, err := GetUserSessionID(c)
	if err != nil {
		log.Printf("[StudentService][HomeStudentPageHandler] error getting user id from session: %+v\n", err)
		return Logout(c)
	}

	student, err := s.uc.GetStudentProfile(id)
	if err != nil {
		log.Printf("[StudentService][HomeStudentPageHandler] error getting user details: %+v\n", err)
		return Logout(c)
	}

	return c.Render(http.StatusOK, "home_student", map[string]interface{}{
		"student": student,
	})
}

func (s *StudentService) ProfileStudentPageHandler(c echo.Context) error {
	id, err := GetUserSessionID(c)
	if err != nil {
		log.Printf("[StudentService][ProfileStudentPageHandler] error getting user id from session: %+v\n", err)
		return Logout(c)
	}

	student, err := s.uc.GetStudentProfile(id)
	if err != nil {
		log.Printf("[StudentService][ProfileStudentPageHandler] error getting user details: %+v\n", err)
		return Logout(c)
	}

	return c.Render(http.StatusOK, "settings", map[string]interface{}{
		"user": student,
		"role": "teacher",
	})
}

func (s *StudentService) RegisteredWebinarHandler(c echo.Context) error {
	id, err := GetUserSessionID(c)
	if err != nil {
		log.Printf("[StudentService][RegisteredWebinarHandler] error getting user id from session: %+v\n", err)
		return BackToHome(c)
	}

	student, err := s.uc.GetStudentProfile(id)
	if err != nil {
		log.Printf("[StudentService][RegisteredWebinarHandler] error getting user details: %+v\n", err)
		return BackToHome(c)
	}

	webinars, err := s.uc.GetAllRegisteredWebinar(id)
	if err != nil {
		log.Printf("[StudentService][RegisteredWebinarHandler] error getting webinar: %+v\n", err)
		return BackToHome(c)
	}

	return c.Render(http.StatusOK, "", map[string]interface{}{
		"student":  student,
		"webinars": webinars,
	})
}

func (s *StudentService) SearchWebinarHandler(c echo.Context) error {
	slug := c.QueryParam("slug")

	webinars, err := s.uc.SearchWebinarBySlug(slug)
	if err != nil {
		log.Printf("[StudentService][SearchWebinarHandler] error getting webinars: %+v\n", err)
		return BackToHome(c)
	}

	return c.Render(http.StatusOK, "webinar_search", map[string]interface{}{
		"webinars": webinars,
		"total_result": len(webinars),
		"slug": slug,
	})
}

func (s *StudentService) StudentWebinarDetailPageHandler(c echo.Context) error {
	return nil
}

func (s *StudentService) EnrollWebinarHandler(c echo.Context) error {
	id, err := GetUserSessionID(c)
	if err != nil {
		log.Printf("[StudentService][EnrollWebinarHandler] error getting user id from session: %+v\n", err)
		return BackToHome(c)
	}

	idParam := c.Param("id")
	webinarID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Printf("[StudentService][EnrollWebinarHandler] error parsing id from query param: %+v\n", err)
		return BackToHome(c)
	}

	webinar, err := s.uc.GetWebinarByID(webinarID)
	if err != nil {
		log.Printf("[StudentService][EnrollWebinarHandler] error getting webinar details: %+v\n", err)
		return BackToHome(c)
	}

	enrollParam := entity.EnrollWebinarParam{
		WebinarID: webinar.ID,
		TeacherID: webinar.TeacherID,
	}

	err = s.uc.EnrollWebinar(id, enrollParam)
	if err != nil {
		log.Printf("[StudentService][EnrollWebinarHandler] error when enrolling student to webinar: %+v\n", err)
		return BackToHome(c)
	}

	return c.Render(http.StatusOK, "", nil)
}

func (s *StudentService) CancelEnrollWebinarHandler(c echo.Context) error {
	id, err := GetUserSessionID(c)
	if err != nil {
		log.Printf("[StudentService][CancelEnrollWebinarHandler] error getting user id from session: %+v\n", err)
		return BackToHome(c)
	}

	idParam := c.Param("id")
	webinarID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Printf("[StudentService][CancelEnrollWebinarHandler] error parsing id from query param: %+v\n", err)
		return BackToHome(c)
	}

	webinar, err := s.uc.GetWebinarByID(webinarID)
	if err != nil {
		log.Printf("[StudentService][CancelEnrollWebinarHandler] error getting webinar details: %+v\n", err)
		return BackToHome(c)
	}

	cancelParam := entity.CancelEnrollmentWebinarParam{
		WebinarID: webinar.ID,
		TeacherID: webinar.TeacherID,
	}

	err = s.uc.CancelEnrollmentWebinar(id, cancelParam)
	if err != nil {
		log.Printf("[StudentService][CancelEnrollWebinarHandler] error when cancelling enrolling student to webinar: %+v\n", err)
		return BackToHome(c)
	}

	return c.Render(http.StatusOK, "", nil)
}
