package service

import (
	"github.com/service/model"
)

type StudentRegisterService struct {
	StudentRegisterModel *model.StudentRegisterModel
}

func NewStudentRegisterService(context *AppContext) (*StudentRegisterService, error) {
	applicationService := &StudentRegisterService{
		StudentRegisterModel: context.Models["studentRegisterModel"].(*model.StudentRegisterModel),
	}
	return applicationService, nil
}
func (s *StudentRegisterService) Register(name, phone, subject, address string) error {
	return s.StudentRegisterModel.Register(name, phone, subject, address)
}
