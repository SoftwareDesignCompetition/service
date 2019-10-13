package service

import (
	"github.com/service/model"
)

type StudentService struct {
	StudentModel *model.StudentModel
}

func NewStudentService(context *AppContext) (*StudentService, error) {
	applicationService := &StudentService{
		StudentModel: context.Models["studentModel"].(*model.StudentModel),
	}
	return applicationService, nil
}
func (s *StudentService) Register(student model.Student) error {
	return s.StudentModel.Register(student)
}

func (s *StudentService) GetAllPhone() ([]string,error) {
	return s.StudentModel.GetAllPhone()
}

func (s *StudentService) GetStudentFormPhone(phone string) (model.Student,error) {
	return s.StudentModel.GetStudentFormPhone(phone)
}

func (s *StudentService) ChangePhone(phone,new_phone string) error {
	return s.StudentModel.ChangePhone(phone,new_phone)
}

func (s *StudentService) ChangeAddress(phone,address string) error {
	return s.StudentModel.ChangePhone(phone,address)
}

func (s *StudentService) ChangeGrade(phone,grade string) error {
	return s.StudentModel.ChangePhone(phone,grade)
}

func (s *StudentService) ChangeSubject(phone,subject string) error {
	return s.StudentModel.ChangePhone(phone,subject)
}

func (s *StudentService) ChangeName(phone,name string) error {
	return s.StudentModel.ChangePhone(phone,name)
}