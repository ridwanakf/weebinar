package service

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ridwanakf/weebinar/internal/entity"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ridwanakf/weebinar/internal"
	"github.com/ridwanakf/weebinar/internal/app"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type CommonService struct {
	teacherUC         internal.TeacherUC
	studentUC         internal.StudentUC
	googleOauthConfig *oauth2.Config
	oAuthRandomState  string
}

func NewCommonService(app *app.WeebinarApp) *CommonService {
	return &CommonService{
		teacherUC: app.UseCases.TeacherUC,
		studentUC: app.UseCases.StudentUC,
		googleOauthConfig: &oauth2.Config{
			ClientID:     app.Cfg.OAuth.ClientID,
			ClientSecret: app.Cfg.OAuth.ClientSecret,
			Endpoint:     google.Endpoint,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
		},
		oAuthRandomState: app.Cfg.OAuth.RandomState,
	}
}

func (s *CommonService) IndexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func (s *CommonService) TeacherSignInHandler(c echo.Context) error {
	return s.oAuthRedirect(c, "teacher")
}

func (s *CommonService) TeacherSignInCallbackHandler(c echo.Context) error {
	var (
		user    entity.GoogleAuth
		teacher entity.Teacher
		role    = "teacher"
	)

	userBytes, err := s.oAuthCallback(c, role)
	if err != nil {
		log.Printf("[CommonService][TeacherSignInCallbackHandler] error executing oAuthCallback: %+v\n", err)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		log.Printf("[CommonService][TeacherSignInCallbackHandler] error unmarshal response content: %+v\n", err)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	teacher.ID, err = strconv.ParseInt(user.ID, 10, 64)
	if err == nil {
		fmt.Printf("[CommonService][TeacherSignInCallbackHandler] error converting id from string to int64: %+v\n", err)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	teacher.Email = user.Email
	teacher.Name = user.Name
	teacher.Picture = user.Picture

	err = s.teacherUC.TeacherSignIn(teacher.ID)
	if err == entity.ErrUserNotFound {
		err = s.teacherUC.TeacherSignUp(teacher)
		if err != nil {
			log.Printf("[CommonService][TeacherSignInCallbackHandler] error creating new user: %+v\n", err)
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}
	}
	s.createSession(c, teacher.ID, teacher.Email, role)

	return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/%s/home", role))
}

func (s *CommonService) StudentSignInHandler(c echo.Context) error {
	return s.oAuthRedirect(c, "student")
}

func (s *CommonService) StudentSignInCallbackHandler(c echo.Context) error {
	var (
		user    entity.GoogleAuth
		student entity.Student
		role    = "student"
	)

	userBytes, err := s.oAuthCallback(c, role)
	if err != nil {
		log.Printf("[CommonService][StudentSignInCallbackHandler] error executing oAuthCallback: %+v\n", err)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		log.Printf("[CommonService][StudentSignInCallbackHandler] error unmarshal response content: %+v\n", err)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	student.ID, err = strconv.ParseInt(user.ID, 10, 64)
	if err == nil {
		fmt.Printf("[CommonService][StudentSignInCallbackHandler] error converting id from string to int64: %+v\n", err)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	student.Email = user.Email
	student.Name = user.Name
	student.Picture = user.Picture

	err = s.studentUC.StudentSignIn(student.ID)
	if err == entity.ErrUserNotFound {
		err = s.studentUC.StudentSignUp(student)
		if err != nil {
			log.Printf("[CommonService][StudentSignInCallbackHandler] error creating new user: %+v\n", err)
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}
	}
	s.createSession(c, student.ID, student.Email, role)

	return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/%s/home", role))
}

func (s *CommonService) oAuthRedirect(c echo.Context, role string) error {
	s.googleOauthConfig.RedirectURL = "http://" + c.Request().Host + "/" + role + "/login/callback"
	url := s.googleOauthConfig.AuthCodeURL(s.oAuthRandomState)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (s *CommonService) oAuthCallback(c echo.Context, role string) ([]byte, error) {
	if c.FormValue("state") != s.oAuthRandomState {
		return nil, errors.New("state is not valid")
	}

	token, err := s.googleOauthConfig.Exchange(oauth2.NoContext, c.FormValue("code"))
	if err != nil {
		return nil, errors.Wrapf(err, "could not get token: ")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get response: ")
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "could not parse response: ")
	}

	return content, nil

	/*
		if role == "teacher" {
			var teacher entity.Teacher
			err = json.Unmarshal(content, &teacher)
			if err != nil {
				log.Printf("[CommonService][oAuthCallback] error unmarshal response content: %+v\n", err)
				return nil, c.Redirect(http.StatusTemporaryRedirect, "/")
			}
			s.createSession(c, teacher.ID, teacher.Email, role)
			return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/%s/home", role))
		} else if role == "student" {
			var student entity.Student
			err = json.Unmarshal(content, &student)
			if err != nil {
				log.Printf("[CommonService][oAuthCallback] error unmarshal response content: %+v\n", err)
				return c.Redirect(http.StatusTemporaryRedirect, "/")
			}
			s.createSession(c, student.ID, student.Email, role)
			return c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/%s/home", role))
		}

		fmt.Fprintf(w, "Response: %s", content)
	*/
}

func (s *CommonService) createSession(c echo.Context, id int64, email string, role string) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 2, // 2 hours
		HttpOnly: true,
	}

	sess.Values["id"] = id
	sess.Values["email"] = email
	sess.Values["role"] = role

	sess.Save(c.Request(), c.Response())
}
