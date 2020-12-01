package service

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ridwanakf/weebinar/internal"
	"github.com/ridwanakf/weebinar/internal/app"
	"github.com/ridwanakf/weebinar/internal/entity"
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
	id, err := GetUserSessionID(c)
	if err != nil {
		log.Printf("[TeacherService][HomeTeacherPageHandler] error getting user id from session: %+v\n", err)
		return Logout(c)
	}

	teacher, err := s.uc.GetTeacherProfile(id)
	if err != nil {
		log.Printf("[TeacherService][HomeTeacherPageHandler] error getting user details: %+v\n", err)
		return Logout(c)
	}

	webinars, err := s.uc.GetAllWebinar(id)
	if err != nil {
		log.Printf("[TeacherService][HomeTeacherPageHandler] error getting webinar: %+v\n", err)
		return Logout(c)
	}

	return c.Render(http.StatusOK, "home_teacher", map[string]interface{}{
		"teacher":  teacher,
		"webinars": webinars,
	})
}

func (s *TeacherService) ProfileTeacherPageHandler(c echo.Context) error {
	id, err := GetUserSessionID(c)
	if err != nil {
		log.Printf("[TeacherService][ProfileTeacherPageHandler] error getting user id from session: %+v\n", err)
		return Logout(c)
	}

	teacher, err := s.uc.GetTeacherProfile(id)
	if err != nil {
		log.Printf("[TeacherService][ProfileTeacherPageHandler] error getting user details: %+v\n", err)
		return Logout(c)
	}
	return c.Render(http.StatusOK, "settings", map[string]interface{}{
		"teacher": teacher,
	})
}

func (s *TeacherService) CreateWebinarPageHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "create", nil)
}

func (s *TeacherService) CreateWebinarPostHandler(c echo.Context) error {
	webinarParam := entity.CreateWebinarParam{}
	if err := c.Bind(&webinarParam); err != nil {
		log.Printf("[TeacherService][CreateWebinarPostHandler] error binding model request: %+v\n", err)
		return BackToHome(c)
	}

	id, err := GetUserSessionID(c)
	if err != nil {
		log.Printf("[TeacherService][CreateWebinarPostHandler] error getting user id from session: %+v\n", err)
		return BackToHome(c)
	}

	webinarParam.CreatedAt = time.Now()

	_, err = s.uc.CreateNewWebinar(id, webinarParam)
	if err != nil {
		log.Printf("[TeacherService][CreateWebinarPostHandler] error creating new webinar: %+v\n", err)
		// TODO: render failed page
		return c.Render(http.StatusOK, "", nil)
	}

	// TODO: render thank you page
	return c.Render(http.StatusOK, "", nil)
}

func (s *TeacherService) WebinarDetailPageHandler(c echo.Context) error {
	idParam := c.Param("id")

	webinarID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Printf("[TeacherService][WebinarDetailPageHandler] error parsing id from query param: %+v\n", err)
		return BackToHome(c)
	}

	webinar, err := s.uc.GetWebinarByID(webinarID)
	if err != nil {
		log.Printf("[TeacherService][WebinarDetailPageHandler] error getting webinar details: %+v\n", err)
		return BackToHome(c)
	}

	return c.Render(http.StatusOK, "detail", map[string]interface{}{
		"webinar": webinar,
	})
}

func (s *TeacherService) UpdateWebinarHandler(c echo.Context) error {
	idParam := c.Param("id")

	webinarID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Printf("[TeacherService][UpdateWebinarHandler] error parsing id from query param: %+v\n", err)
		return BackToHome(c)
	}

	webinarParam := entity.UpdateWebinarParam{}
	if err := c.Bind(&webinarParam); err != nil {
		log.Printf("[TeacherService][UpdateWebinarHandler] error binding model request: %+v\n", err)
		return BackToHome(c)
	}
	webinarParam.ID = webinarID

	teacherID, err := GetUserSessionID(c)
	if err != nil {
		log.Printf("[TeacherService][UpdateWebinarHandler] error getting user id from session: %+v\n", err)
		return Logout(c)
	}

	err = s.uc.UpdateWebinar(teacherID, webinarParam)
	if err != nil {
		log.Printf("[TeacherService][UpdateWebinarHandler] error when updating webinar: %+v\n", err)
		return BackToHome(c)
	}

	// TODO: render popup message
	return BackToHome(c)
}

func (s *TeacherService) DeleteWebinarHandler(c echo.Context) error {
	idParam := c.Param("id")

	webinarID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Printf("[TeacherService][DeleteWebinarHandler] error parsing id from query param: %+v\n", err)
		return BackToHome(c)
	}
	webinarParam := entity.DeleteWebinarParam{ID: webinarID}

	teacherID, err := GetUserSessionID(c)
	if err != nil {
		log.Printf("[TeacherService][DeleteWebinarHandler] error getting user id from session: %+v\n", err)
		return Logout(c)
	}

	err = s.uc.DeleteWebinar(teacherID, webinarParam)
	if err != nil {
		log.Printf("[TeacherService][DeleteWebinarHandler] error when updating webinar: %+v\n", err)
		return BackToHome(c)
	}

	// TODO: render popup message
	return BackToHome(c)
}
