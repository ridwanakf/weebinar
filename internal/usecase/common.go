package usecase

import "github.com/ridwanakf/weebinar/internal"

type CommonUsecase struct {
	repo internal.CommonRepo
}

func NewCommonUsecase(repo internal.CommonRepo) *CommonUsecase {
	return &CommonUsecase{repo: repo}
}

func (c *CommonUsecase) CheckUser(id int64) error {
	panic("implement me")
}

func (c *CommonUsecase) TeacherSignUp() error {
	panic("implement me")
}

func (c *CommonUsecase) TeacherSignIn() error {
	panic("implement me")
}

func (c *CommonUsecase) StudentSignUp() error {
	panic("implement me")
}

func (c *CommonUsecase) StudentSignIn() error {
	panic("implement me")
}
