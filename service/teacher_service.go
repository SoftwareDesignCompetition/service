package service

import (
	"github.com/service/model"
)

type TeacherService struct {
	TeacherModel *model.TeacherModel
}

func NewTeacherService(context *AppContext) (*TeacherService, error) {
	return &TeacherService{
		TeacherModel: context.Models["teacherModel"].(*model.TeacherModel),
	}, nil
}

func (s *TeacherService) Register(teacher model.Teacher) error {
	return s.TeacherModel.Register(teacher)
}

func (s *TeacherService) ChangePhone(phone, new_phone string) error {
	return s.TeacherModel.ChangePhone(phone, new_phone)
}

func (s *TeacherService) ChangeAddress(phone, address string) error {
	return s.TeacherModel.ChangeAddress(phone, address)
}

func (s *TeacherService) ChangeGrade(phone, grade string) error {
	return s.TeacherModel.ChangeGrade(phone, grade)
}

func (s *TeacherService) ChangeSubject(phone, subject string) error {
	return s.TeacherModel.ChangeSubject(phone, subject)
}

func (s *TeacherService) ChangeTeGrade(phone, tegrade string) error {
	return s.TeacherModel.ChangeTeGrade(phone, tegrade)
}

func (s *TeacherService) ChangeName(phone, name string) error {
	return s.TeacherModel.ChangeName(phone, name)
}

func (s *TeacherService) ChangeSchool(phone, school string) error {
	return s.TeacherModel.ChangeSchool(phone, school)
}

func (s *TeacherService) ChangeGender(phone, gender string) error {
	return s.TeacherModel.ChangeGender(phone, gender)
}

func (s *TeacherService) ChangeTimes(phone, times string) error {
	return s.TeacherModel.ChangeTimes(phone, times)
}

func (s *TeacherService) ChangeSalary(phone, salary string) error {
	return s.TeacherModel.ChangeSalary(phone, salary)
}

func (s *TeacherService) ChangeMajor(phone, major string) error {
	return s.TeacherModel.ChangeMajor(phone, major)
}

func (s *TeacherService) ChangeStatus(phone string) error {
	return s.TeacherModel.ChangeStatus(phone)
}

func (s *TeacherService) AddEvaluate(phone string, evaluate bool) error {
	return s.TeacherModel.AddEvaluate(phone, evaluate)
}

func (s *TeacherService) GetAllPhone() ([]string, error) {
	return s.TeacherModel.GetAllPhone()
}


func (s *TeacherService) GetTeacher(newGrade, newSubject, newMoney, newSex string) ([]model.Teacher, error) {
	return s.TeacherModel.GetTeacher(newGrade, newSubject, newMoney, newSex)
}