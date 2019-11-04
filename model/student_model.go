package model

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/service/config"
	"github.com/service/rediskey"
	"github.com/tb_common/cache"
	"strconv"
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
	_, err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+student.Phone+":good", 0).Result()
	if err != nil {
		log.Errorf("add grade err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+student.Phone+":own", 0).Result()
	if err != nil {
		log.Errorf("add grade err, %v", err)
		return err
	}
	return nil
}

func (s *StudentModel) GetAllPhone() ([]string, error) {
	req, err := s.GetRedisMaster().HGetAll(rediskey.GetStudentRedisKey()).Result()
	if err != nil {
		log.Errorf("get allphone err, %v", err)
		return nil, err
	}
	var resp []string
	if req == nil {
		return nil, nil
	}
	for i, _ := range req {
		phone := i[5:16]
		resp = append(resp, phone)
	}
	resp = s.RemoveRepByMap(resp)
	return resp, nil
}

func (s *StudentModel) RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

func (s *StudentModel) GetStudentFormPhone(phone string) (Student, error) {
	name, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(), "user:"+phone+":name").Result()
	if name == "" {
		return Student{}, nil
	}
	if err != nil {
		log.Errorf("get name err, %v", err)
		return Student{}, err
	}
	subject, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(), "user:"+phone+":subject").Result()
	if err != nil {
		log.Errorf("get subject err, %v", err)
		return Student{}, err
	}
	address, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(), "user:"+phone+":address").Result()
	if err != nil {
		log.Errorf("get address err, %v", err)
		return Student{}, err
	}
	grade, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(), "user:"+phone+":grade").Result()
	if err != nil {
		log.Errorf("get grade err, %v", err)
		return Student{}, err
	}
	good, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(), "user:"+phone+":good").Result()
	if err != nil {
		log.Errorf("get good err, %v", err)
		return Student{}, err
	}
	good1, err := strconv.Atoi(good)
	if err != nil {
		log.Errorf("change good err, %v", err)
		return Student{}, err
	}
	own, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(), "user:"+phone+":own").Result()
	if err != nil {
		log.Errorf("get own err, %v", err)
		return Student{}, err
	}
	own1, err := strconv.Atoi(own)
	if err != nil {
		log.Errorf("change own err, %v", err)
		return Student{}, err
	}
	return Student{
		Phone:   phone,
		Name:    name,
		Address: address,
		Subject: subject,
		Grade:   grade,
		Good:    good1,
		Own:     own1,
	}, nil
}

func (s *StudentModel) ChangePhone(phone, new_phone string) error {
	address, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(), "user:"+phone+":address").Result()
	if err != nil {
		log.Errorf("get address err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+new_phone+":address", address).Result()
	if err != nil {
		log.Errorf("set address err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetStudentRedisKey(), "user:"+phone+":address").Result()
	if err != nil {
		log.Errorf("del address err, %v", err)
		return err
	}
	name, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(), "user:"+phone+":name").Result()
	if err != nil {
		log.Errorf("get name err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+new_phone+":name", name).Result()
	if err != nil {
		log.Errorf("set name err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetStudentRedisKey(), "user:"+phone+":name").Result()
	if err != nil {
		log.Errorf("del subject err, %v", err)
		return err
	}
	subject, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(), "user:"+phone+":subject").Result()
	if err != nil {
		log.Errorf("get subject err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+new_phone+":subject", subject).Result()
	if err != nil {
		log.Errorf("set subject err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetStudentRedisKey(), "user:"+phone+":subject").Result()
	if err != nil {
		log.Errorf("del subject err, %v", err)
		return err
	}
	grade, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(), "user:"+phone+":grade").Result()
	if err != nil {
		log.Errorf("get grade err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+new_phone+":grade", grade).Result()
	if err != nil {
		log.Errorf("set grade err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetStudentRedisKey(), "user:"+phone+":grade").Result()
	if err != nil {
		log.Errorf("del grade err, %v", err)
		return err
	}
	return nil
}

func (s *StudentModel) ChangeAddress(phone, address string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+phone+":address", address).Result()
	if err != nil {
		log.Errorf("change address err, %v", err)
		return err
	}
	return nil
}

func (s *StudentModel) ChangeGrade(phone, grade string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+phone+":grade", grade).Result()
	if err != nil {
		log.Errorf("change grade err, %v", err)
		return err
	}
	return nil
}

func (s *StudentModel) ChangeSubject(phone, subject string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+phone+":subject", subject).Result()
	if err != nil {
		log.Errorf("change subject err, %v", err)
		return err
	}
	return nil
}

func (s *StudentModel) ChangeName(phone, name string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+phone+":name", name).Result()
	if err != nil {
		log.Errorf("change name err, %v", err)
		return err
	}
	return nil
}

func (s *StudentModel) AddEvaluate(phone string, evaluate bool) error {
	own, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(), "user:"+phone+":own").Result()
	if err != nil {
		log.Errorf("get own err, %v", err)
		return err
	}
	own1, err := strconv.Atoi(own)
	if err != nil {
		log.Errorf("change own err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+phone+":own", fmt.Sprint(own1+1)).Result()
	if err != nil {
		log.Errorf("set own err, %v", err)
		return err
	}
	if evaluate {
		good, err := s.GetRedisMaster().HGet(rediskey.GetStudentRedisKey(), "user:"+phone+":good").Result()
		if err != nil {
			log.Errorf("get good err, %v", err)
			return err
		}
		good1, err := strconv.Atoi(good)
		if err != nil {
			log.Errorf("change good err, %v", err)
			return err
		}
		_, err = s.GetRedisMaster().HSet(rediskey.GetStudentRedisKey(), "user:"+phone+":good", fmt.Sprint(good1+1)).Result()
		if err != nil {
			log.Errorf("set good err, %v", err)
			return err
		}
	}
	return nil
}
