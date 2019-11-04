package model

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/service/config"
	"github.com/service/rediskey"
	"github.com/tb_common/cache"
	"strconv"
)

type TeacherModel struct {
	BaseCacheModel
}

func NewTeacherModel(config *config.AppConfig) (*TeacherModel, error) {
	caches, err := cache.GetCaches(&config.Redis)
	if err != nil {
		log.Errorf("get caches error: %v", err)
		return nil, err
	}
	return &TeacherModel{
		BaseCacheModel: BaseCacheModel{
			Config: config,
			Dbs:    caches,
		},
	}, nil
}

func (s *TeacherModel) Register(teacher Teacher) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+teacher.Phone+":name", teacher.Name).Result()
	if err != nil {
		log.Errorf("add name err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+teacher.Phone+":subject", teacher.Subject).Result()
	if err != nil {
		log.Errorf("add subject err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+teacher.Phone+":address", teacher.Address).Result()
	if err != nil {
		log.Errorf("add address err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+teacher.Phone+":grade", teacher.Grade).Result()
	if err != nil {
		log.Errorf("add grade err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+teacher.Phone+":school", teacher.School).Result()
	if err != nil {
		log.Errorf("add school err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+teacher.Phone+":gender", teacher.Gender).Result()
	if err != nil {
		log.Errorf("add gender err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+teacher.Phone+":times", teacher.Times).Result()
	if err != nil {
		log.Errorf("add times err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+teacher.Phone+":salary", teacher.Salary).Result()
	if err != nil {
		log.Errorf("add salary err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+teacher.Phone+":major", teacher.Major).Result()
	if err != nil {
		log.Errorf("add major err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+teacher.Phone+":status", 0).Result()
	if err != nil {
		log.Errorf("add status err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+teacher.Phone+":good", 0).Result()
	if err != nil {
		log.Errorf("add good err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+teacher.Phone+":own", 0).Result()
	if err != nil {
		log.Errorf("add own err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+teacher.Phone+":teGrade", teacher.TeGrade).Result()
	if err != nil {
		log.Errorf("add tegrade err, %v", err)
		return err
	}
	return nil
}

func (s *TeacherModel) ChangePhone(phone, new_phone string) error {
	name, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":name").Result()
	if err != nil {
		log.Errorf("get name err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+new_phone+":name", name).Result()
	if err != nil {
		log.Errorf("set name err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetTeacherRedisKey(), "user:"+phone+":name").Result()
	if err != nil {
		log.Errorf("del name err, %v", err)
		return err
	}
	subject, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":subject").Result()
	if err != nil {
		log.Errorf("get subject err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+new_phone+":subject", subject).Result()
	if err != nil {
		log.Errorf("set subject err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetTeacherRedisKey(), "user:"+phone+":subject").Result()
	if err != nil {
		log.Errorf("del subject err, %v", err)
		return err
	}
	address, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":address").Result()
	if err != nil {
		log.Errorf("get address err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+new_phone+":address", address).Result()
	if err != nil {
		log.Errorf("set address err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetTeacherRedisKey(), "user:"+phone+":address").Result()
	if err != nil {
		log.Errorf("del address err, %v", err)
		return err
	}
	tegrade, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":tegrade").Result()
	if err != nil {
		log.Errorf("get tegrade err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+new_phone+":tegrade", tegrade).Result()
	if err != nil {
		log.Errorf("set tegrade err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetTeacherRedisKey(), "user:"+phone+":tegrade").Result()
	if err != nil {
		log.Errorf("del tegrade err, %v", err)
		return err
	}
	grade, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":grade").Result()
	if err != nil {
		log.Errorf("get grade err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+new_phone+":grade", grade).Result()
	if err != nil {
		log.Errorf("set grade err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetTeacherRedisKey(), "user:"+phone+":grade").Result()
	if err != nil {
		log.Errorf("del grade err, %v", err)
		return err
	}
	school, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":school").Result()
	if err != nil {
		log.Errorf("get school err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+new_phone+":school", school).Result()
	if err != nil {
		log.Errorf("set school err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetTeacherRedisKey(), "user:"+phone+":school").Result()
	if err != nil {
		log.Errorf("del school err, %v", err)
		return err
	}
	gender, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":gender").Result()
	if err != nil {
		log.Errorf("get gender err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+new_phone+":gender", gender).Result()
	if err != nil {
		log.Errorf("set gender err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetTeacherRedisKey(), "user:"+phone+":gender").Result()
	if err != nil {
		log.Errorf("del gender err, %v", err)
		return err
	}
	times, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":times").Result()
	if err != nil {
		log.Errorf("get times err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+new_phone+":times", times).Result()
	if err != nil {
		log.Errorf("set times err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetTeacherRedisKey(), "user:"+phone+":times").Result()
	if err != nil {
		log.Errorf("del times err, %v", err)
		return err
	}
	salary, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":salary").Result()
	if err != nil {
		log.Errorf("get salary err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+new_phone+":salary", salary).Result()
	if err != nil {
		log.Errorf("set salary err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetTeacherRedisKey(), "user:"+phone+":salary").Result()
	if err != nil {
		log.Errorf("del salary err, %v", err)
		return err
	}
	major, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":major").Result()
	if err != nil {
		log.Errorf("get major err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+new_phone+":major", major).Result()
	if err != nil {
		log.Errorf("set major err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetTeacherRedisKey(), "user:"+phone+":major").Result()
	if err != nil {
		log.Errorf("del major err, %v", err)
		return err
	}
	status, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":status").Result()
	if err != nil {
		log.Errorf("get status err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+new_phone+":status", status).Result()
	if err != nil {
		log.Errorf("set status err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetTeacherRedisKey(), "user:"+phone+":status").Result()
	if err != nil {
		log.Errorf("del status err, %v", err)
		return err
	}
	good, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":good").Result()
	if err != nil {
		log.Errorf("get good err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+new_phone+":good", good).Result()
	if err != nil {
		log.Errorf("set good err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetTeacherRedisKey(), "user:"+phone+":good").Result()
	if err != nil {
		log.Errorf("del good err, %v", err)
		return err
	}
	own, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":own").Result()
	if err != nil {
		log.Errorf("get own err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+new_phone+":own", own).Result()
	if err != nil {
		log.Errorf("set own err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HDel(rediskey.GetTeacherRedisKey(), "user:"+phone+":own").Result()
	if err != nil {
		log.Errorf("del own err, %v", err)
		return err
	}
	return nil
}

func (s *TeacherModel) ChangeAddress(phone, address string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+phone+":address", address).Result()
	if err != nil {
		log.Errorf("change address err, %v", err)
		return err
	}
	return nil
}

func (s *TeacherModel) ChangeGrade(phone, grade string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+phone+":grade", grade).Result()
	if err != nil {
		log.Errorf("change grade err, %v", err)
		return err
	}
	return nil
}

func (s *TeacherModel) ChangeSubject(phone, subject string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+phone+":subject", subject).Result()
	if err != nil {
		log.Errorf("change subject err, %v", err)
		return err
	}
	return nil
}

func (s *TeacherModel) ChangeTeGrade(phone, tegrade string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+phone+":tegrade", tegrade).Result()
	if err != nil {
		log.Errorf("change tegrade err, %v", err)
		return err
	}
	return nil
}

func (s *TeacherModel) ChangeName(phone, name string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+phone+":name", name).Result()
	if err != nil {
		log.Errorf("change name err, %v", err)
		return err
	}
	return nil
}

func (s *TeacherModel) ChangeSchool(phone, school string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+phone+":school", school).Result()
	if err != nil {
		log.Errorf("change school err, %v", err)
		return err
	}
	return nil
}

func (s *TeacherModel) ChangeGender(phone, gender string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+phone+":gender", gender).Result()
	if err != nil {
		log.Errorf("change gender err, %v", err)
		return err
	}
	return nil
}

func (s *TeacherModel) ChangeTimes(phone, times string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+phone+":times", times).Result()
	if err != nil {
		log.Errorf("change times err, %v", err)
		return err
	}
	return nil
}

func (s *TeacherModel) ChangeSalary(phone, salary string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+phone+":salary", salary).Result()
	if err != nil {
		log.Errorf("change salary err, %v", err)
		return err
	}
	return nil
}

func (s *TeacherModel) ChangeMajor(phone, major string) error {
	_, err := s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+phone+":major", major).Result()
	if err != nil {
		log.Errorf("change major err, %v", err)
		return err
	}
	return nil
}

func (s *TeacherModel) ChangeStatus(phone string) error {
	status, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":status").Result()
	if err != nil {
		log.Errorf("get status err, %v", err)
		return err
	}
	status1, err := strconv.Atoi(status)
	if err != nil {
		log.Errorf("change status err, %v", err)
		return err
	}
	if status1 == 0 {
		_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+phone+":status", 1).Result()
	} else {
		_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+phone+":status", 0).Result()
	}
	if err != nil {
		log.Errorf("set status err, %v", err)
		return err
	}
	return nil
}

func (s *TeacherModel) AddEvaluate(phone string, evaluate bool) error {
	own, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":own").Result()
	if err != nil {
		log.Errorf("get own err, %v", err)
		return err
	}
	own1, err := strconv.Atoi(own)
	if err != nil {
		log.Errorf("change own err, %v", err)
		return err
	}
	_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+phone+":own", fmt.Sprint(own1+1)).Result()
	if err != nil {
		log.Errorf("set own err, %v", err)
		return err
	}
	if evaluate {
		good, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+phone+":good").Result()
		if err != nil {
			log.Errorf("get good err, %v", err)
			return err
		}
		good1, err := strconv.Atoi(good)
		if err != nil {
			log.Errorf("change good err, %v", err)
			return err
		}
		_, err = s.GetRedisMaster().HSet(rediskey.GetTeacherRedisKey(), "user:"+phone+":good", fmt.Sprint(good1+1)).Result()
		if err != nil {
			log.Errorf("set good err, %v", err)
			return err
		}
	}
	return nil
}

func (s *TeacherModel) GetAllPhone() ([]string, error) {
	req, err := s.GetRedisMaster().HGetAll(rediskey.GetTeacherRedisKey()).Result()
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

func (s *TeacherModel) RemoveRepByMap(slc []string) []string {
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

func (s *TeacherModel) GetTeacher(newGrade, newSubject, newMoney, newSex string) ([]Teacher, error) {
	req, err := s.GetAllPhone()
	if err != nil {
		log.Errorf("get GetAllPhone err, %v", err)
		return nil, err
	}
	if req == nil {
		return nil, nil
	}
	var resp []Teacher

	for _, i := range req {
		name, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+i+":name").Result()
		if err != nil {
			log.Errorf("get name err, %v", err)
			return nil, err
		}
		subject, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+i+":subject").Result()
		if err != nil {
			log.Errorf("get subject err, %v", err)
			return nil, err
		}
		teGrade, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+i+":teGrade").Result()
		if err != nil {
			log.Errorf("get teGrade err, %v", err)
			return nil, err
		}
		address, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+i+":address").Result()
		if err != nil {
			log.Errorf("get address err, %v", err)
			return nil, err
		}
		grade, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+i+":grade").Result()
		if err != nil {
			log.Errorf("get grade err, %v", err)
			return nil, err
		}
		school, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+i+":school").Result()
		if err != nil {
			log.Errorf("get school err, %v", err)
			return nil, err
		}
		status, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+i+":status").Result()
		if err != nil {
			log.Errorf("get status err, %v", err)
			return nil, err
		}
		gender, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+i+":gender").Result()
		if err != nil {
			log.Errorf("get gender err, %v", err)
			return nil, err
		}
		times, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+i+":times").Result()
		if err != nil {
			log.Errorf("get times err, %v", err)
			return nil, err
		}
		salary, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+i+":salary").Result()
		if err != nil {
			log.Errorf("get salary err, %v", err)
			return nil, err
		}
		major, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+i+":major").Result()
		if err != nil {
			log.Errorf("get major err, %v", err)
			return nil, err
		}
		good, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+i+":good").Result()
		if err != nil {
			log.Errorf("get good err, %v", err)
			return nil, err
		}
		own, err := s.GetRedisMaster().HGet(rediskey.GetTeacherRedisKey(), "user:"+i+":own").Result()
		if err != nil {
			log.Errorf("get status err, %v", err)
			return nil, err
		}
		status1, err := strconv.Atoi(status)
		if err != nil {
			log.Errorf("change string err, %v", err)
			return nil, err
		}
		good1, err := strconv.Atoi(good)
		if err != nil {
			log.Errorf("change string err, %v", err)
			return nil, err
		}
		own1, err := strconv.Atoi(own)
		if err != nil {
			log.Errorf("change string err, %v", err)
			return nil, err
		}
		gender1, err := strconv.Atoi(gender)
		if err != nil {
			log.Errorf("change string err, %v", err)
			return nil, err
		}
		salary1, err := strconv.Atoi(salary)
		if err != nil {
			log.Errorf("change string err, %v", err)
			return nil, err
		}
		subject1, err := strconv.Atoi(subject)
		if err != nil {
			log.Errorf("change string err, %v", err)
			return nil, err
		}
		teGrade1, err := strconv.Atoi(teGrade)
		if err != nil {
			log.Errorf("change string err, %v", err)
			return nil, err
		}
		newSubject, err := strconv.Atoi(newSubject)
		if err != nil {
			log.Errorf("change string err, %v", err)
			return nil, err
		}
		newGrade, err := strconv.Atoi(newGrade)
		if err != nil {
			log.Errorf("change string err, %v", err)
			return nil, err
		}
		newMoney, err := strconv.Atoi(newMoney)
		if err != nil {
			log.Errorf("change string err, %v", err)
			return nil, err
		}
		newSex, err := strconv.Atoi(newSex)
		if err != nil {
			log.Errorf("change string err, %v", err)
			return nil, err
		}
		if s.gcd(subject1, newSubject) == 1 {
			continue
		}
		if s.gcd(teGrade1, newGrade) == 1 {
			continue
		}
		if salary1 > newMoney {
			continue
		}
		if newSex != 2 && gender1 != newSex {
			continue
		}
		resp = append(resp, Teacher{
			Phone:   i,       //手机号
			Name:    name,    //姓名
			Subject: subject, //教授科目
			TeGrade: teGrade, //教授年级
			Address: address, //住址
			Grade:   grade,   //教师年级
			School:  school,  //教师学校
			Gender:  gender,  //性别
			Times:   times,   //时间
			Salary:  salary,  //薪资
			Major:   major,   //专业
			Status:  status1, //是否接单
			Good:    good1,   //好评次数
			Own:     own1,    //总的接单次数
		})
	}
	return resp, nil
}

func (s *TeacherModel) gcd(x, y int) int { //计算最大公约数
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
