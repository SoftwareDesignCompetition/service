package model

import (
	log "github.com/cihub/seelog"
	"github.com/service/config"
	"github.com/service/rediskey"
	"github.com/tb_common/cache"
)

type StudentModel struct {
	BaseCacheModel
}

func NewStudentModel(config *config.AppConfig) (*StudentModel, error) {
	caches, err := cache.GetCaches(&config.Redis)
	if err != nil {
		log.Errorf("get caches error: %v", err)
		return nil, err
	}
	return &StudentModel{
		BaseCacheModel: BaseCacheModel{
			Config: config,
			Dbs:    caches,
		},
	}, nil
}

func (s *StudentModel) Register(student Student) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+student.Phone+":name", student.Name).Result()
	if err != nil {
		log.Errorf("add name err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+student.Phone+":subject", student.Subject).Result()
	if err != nil {
		log.Errorf("add subject err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+student.Phone+":address", student.Address).Result()
	if err != nil {
		log.Errorf("add address err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+student.Phone+":grade", student.Grade).Result()
	if err != nil {
		log.Errorf("add grade err, %v", err)
		return err
	}
	return nil
}

func (s *StudentModel) GetAllPhone() ([]string,error) {
	req, err := s.GetRedisMaster().HGetAll(rediskey.GetStudentRedisKey()).Result()
	if err != nil {
		log.Errorf("get allphone err, %v", err)
		return nil,err
	}
	var resp []string
	for i,_ := range req{
		phone := i[5:16]
		resp = append(resp,phone )
	}
	return resp,nil
}

func (s *StudentModel) GetStudentFormPhone(phone string) (Student,error) {
	name, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(),"user:"+phone+":name").Result()
	if err != nil {
		log.Errorf("get name err, %v", err)
		return Student{},err
	}
	subject, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(),"user:"+phone+":subject").Result()
	if err != nil {
		log.Errorf("get subject err, %v", err)
		return Student{},err
	}
	address, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(),"user:"+phone+":address").Result()
	if err != nil {
		log.Errorf("get address err, %v", err)
		return Student{},err
	}
	grade, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(),"user:"+phone+":grade").Result()
	if err != nil {
		log.Errorf("get grade err, %v", err)
		return Student{},err
	}
	return Student{
		Phone:phone,
		Name:name,
		Address:address,
		Subject:subject,
		Grade:grade,
	},nil
}

func (s *StudentModel) ChangePhone(phone,new_phone string) error {
	address,err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(),"user:"+phone+":address").Result()
	if err != nil {
		log.Errorf("get address err, %v", err)
		return err
	}
	_,err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(),"user:"+new_phone+":address",address).Result()
	if err != nil {
		log.Errorf("set address err, %v", err)
		return err
	}
	_,err = s.GetRedisMaster().HDel(rediskey.GetStudentRedisKey(),"user:"+phone+":address").Result()
	if err != nil {
		log.Errorf("del address err, %v", err)
		return err
	}
	name,err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(),"user:"+phone+":name").Result()
	if err != nil {
		log.Errorf("get name err, %v", err)
		return err
	}
	_,err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(),"user:"+new_phone+":name",name).Result()
	if err != nil {
		log.Errorf("set name err, %v", err)
		return err
	}
	_,err = s.GetRedisMaster().HDel(rediskey.GetStudentRedisKey(),"user:"+phone+":name").Result()
	if err != nil {
		log.Errorf("del subject err, %v", err)
		return err
	}
	subject,err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(),"user:"+phone+":subject").Result()
	if err != nil {
		log.Errorf("get subject err, %v", err)
		return err
	}
	_,err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(),"user:"+new_phone+":subject",subject).Result()
	if err != nil {
		log.Errorf("set subject err, %v", err)
		return err
	}
	_,err = s.GetRedisMaster().HDel(rediskey.GetStudentRedisKey(),"user:"+phone+":subject").Result()
	if err != nil {
		log.Errorf("del subject err, %v", err)
		return err
	}
	grade,err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(),"user:"+phone+":grade").Result()
	if err != nil {
		log.Errorf("get grade err, %v", err)
		return err
	}
	_,err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(),"user:"+new_phone+":grade",grade).Result()
	if err != nil {
		log.Errorf("set grade err, %v", err)
		return err
	}
	_,err = s.GetRedisMaster().HDel(rediskey.GetStudentRedisKey(),"user:"+phone+":grade").Result()
	if err != nil {
		log.Errorf("del grade err, %v", err)
		return err
	}
	return nil
}

func (s *StudentModel) ChangeAddress(phone,address string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+phone+":address", address).Result()
	if err != nil {
		log.Errorf("change address err, %v", err)
		return err
	}
	return nil
}

func (s *StudentModel) ChangeGrade(phone,grade string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+phone+":grade", grade).Result()
	if err != nil {
		log.Errorf("change grade err, %v", err)
		return err
	}
	return nil
}

func (s *StudentModel) ChangeSubject(phone,subject string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+phone+":subject", subject).Result()
	if err != nil {
		log.Errorf("change subject err, %v", err)
		return err
	}
	return nil
}

func (s *StudentModel) ChangeName(phone,name string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+phone+":name", name).Result()
	if err != nil {
		log.Errorf("change name err, %v", err)
		return err
	}
	return nil
}