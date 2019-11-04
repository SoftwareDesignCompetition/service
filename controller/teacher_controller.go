package controller

import (
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/service/middleware"
	"github.com/service/model"
	"github.com/service/service"
	"strconv"
	"time"
)

type TeacherController struct {
	BaseController
	TeacherService *service.TeacherService
}

func NewTeacherController(context *service.AppContext) (*TeacherController, error) {
	teacherService := context.Services["teacherService"].(*service.TeacherService)
	return &TeacherController{
		BaseController: BaseController{
			Config: context.Config,
		},
		TeacherService: teacherService,
	}, nil
}

func (s *TeacherController) Register(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone
	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	app_key := c.PostForm("appkey")
	if app_key == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[app_key]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	name := c.PostForm("name")
	if name == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "name is empty",
		})
		return
	}
	param["name"] = name
	subject := c.PostForm("subject")
	if subject == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "subject is empty",
		})
		return
	}
	param["subject"] = subject
	teGrade := c.PostForm("teGrade")
	if teGrade == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "teGrade is empty",
		})
		return
	}
	param["teGrade"] = teGrade
	address := c.PostForm("address")
	if address == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "address is empty",
		})
		return
	}
	param["address"] = address
	grade := c.PostForm("grade")
	if grade == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "grade is empty",
		})
		return
	}
	param["grade"] = grade
	school := c.PostForm("school")
	if school == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "school is empty",
		})
		return
	}
	param["school"] = school
	gender := c.PostForm("gender")
	if gender == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "gender is empty",
		})
		return
	}
	param["gender"] = gender
	times1 := c.PostForm("times")
	if times1 == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "times is empty",
		})
		return
	}
	param["times"] = times1
	salary := c.PostForm("salary")
	if salary == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "salary is empty",
		})
		return
	}
	param["salary"] = salary
	major := c.PostForm("major")
	if major == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "major is empty",
		})
		return
	}
	param["major"] = major
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	teacher := model.Teacher{
		Phone:   phone,   //手机号
		Name:    name,    //姓名
		Subject: subject, //教授科目
		TeGrade: teGrade, //教授年级
		Address: address, //住址
		Grade:   grade,   //教师年级
		School:  school,  //教师学校
		Gender:  gender,  //性别
		Times:   times1,  //时间
		Salary:  salary,  //薪资
		Major:   major,   //专业
	}
	resp, err := s.TeacherService.GetAllPhone() //检验是否已经注册
	for _, i := range resp {
		if i == teacher.Phone {
			c.JSON(201, gin.H{
				"result":  1,
				"message": "Illegal double registration", //非法双重注册
			})
			return
		}
	}
	err = s.TeacherService.Register(teacher)
	if err != nil {
		log.Errorf("teacherRegister error %v", err)
		c.JSON(400, gin.H{
			"result":  2,
			"message": "teacherRegister error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *TeacherController) ChangePhone(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone

	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	app_key := c.PostForm("appkey")
	if app_key == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[app_key]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	new_phone := c.PostForm("new_phone")
	if new_phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "new_phone is empty",
		})
		return
	}
	param["new_phone"] = new_phone
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	err = s.TeacherService.ChangePhone(phone, new_phone)
	if err != nil {
		log.Error("change phone error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change phone error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *TeacherController) ChangeAddress(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone
	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	app_key := c.PostForm("appkey")
	if app_key == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[app_key]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	address := c.PostForm("address")
	if address == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "address is empty",
		})
		return
	}
	param["address"] = address
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	err = s.TeacherService.ChangeAddress(phone, address)
	if err != nil {
		log.Error("change address error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change address error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *TeacherController) ChangeGrade(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone
	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	app_key := c.PostForm("appkey")
	if app_key == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[app_key]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	grade := c.PostForm("grade")
	if grade == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "grade is empty",
		})
		return
	}
	param["grade"] = grade
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	err = s.TeacherService.ChangeGrade(phone, grade)
	if err != nil {
		log.Error("change grade error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change grade error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *TeacherController) ChangeSubject(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone
	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	app_key := c.PostForm("appkey")
	if app_key == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[app_key]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	subject := c.PostForm("subject")
	if subject == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "subject is empty",
		})
		return
	}
	param["subject"] = subject
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	err = s.TeacherService.ChangeSubject(phone, subject)
	if err != nil {
		log.Error("change subject error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change subject error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *TeacherController) ChangeTeGrade(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone
	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	app_key := c.PostForm("appkey")
	if app_key == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[app_key]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	teGrade := c.PostForm("teGrade")
	if teGrade == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "teGrade is empty",
		})
		return
	}
	param["teGrade"] = teGrade
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	err = s.TeacherService.ChangeTeGrade(phone, teGrade)
	if err != nil {
		log.Error("change teGrade error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change teGrade error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *TeacherController) ChangeName(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone
	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	app_key := c.PostForm("appkey")
	if app_key == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[app_key]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	name := c.PostForm("name")
	if name == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "name is empty",
		})
		return
	}
	param["name"] = name
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	err = s.TeacherService.ChangeName(phone, name)
	if err != nil {
		log.Error("change name error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change name error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *TeacherController) ChangeSchool(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone
	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	app_key := c.PostForm("appkey")
	if app_key == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[app_key]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	school := c.PostForm("school")
	if school == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "school is empty",
		})
		return
	}
	param["school"] = school
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	err = s.TeacherService.ChangeSchool(phone, school)
	if err != nil {
		log.Error("change school error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change school error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *TeacherController) ChangeGender(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone
	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	app_key := c.PostForm("appkey")
	if app_key == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[app_key]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	gender := c.PostForm("gender")
	if gender == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "gender is empty",
		})
		return
	}
	param["gender"] = gender
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	err = s.TeacherService.ChangeGender(phone, gender)
	if err != nil {
		log.Error("change gender error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change gender error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *TeacherController) ChangeTimes(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone
	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	app_key := c.PostForm("appkey")
	if app_key == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[app_key]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	times1 := c.PostForm("times")
	if times1 == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "times is empty",
		})
		return
	}
	param["times"] = times
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	err = s.TeacherService.ChangeTimes(phone, times1)
	if err != nil {
		log.Error("change times error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change times error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *TeacherController) ChangeSalary(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone
	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	app_key := c.PostForm("appkey")
	if app_key == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[app_key]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	salary := c.PostForm("salary")
	if salary == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "salary is empty",
		})
		return
	}
	param["salary"] = salary
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	err = s.TeacherService.ChangeSalary(phone, salary)
	if err != nil {
		log.Error("change salary error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change salary error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *TeacherController) ChangeMajor(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone
	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	appKey := c.PostForm("appkey")
	if appKey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[appKey]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	major := c.PostForm("major")
	if major == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "major is empty",
		})
		return
	}
	param["major"] = major
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	err = s.TeacherService.ChangeMajor(phone, major)
	if err != nil {
		log.Error("change major error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change major error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *TeacherController) ChangeStatus(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone
	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	appKey := c.PostForm("appkey")
	if appKey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[appKey]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	err = s.TeacherService.ChangeStatus(phone)
	if err != nil {
		log.Error("change status error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change status error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *TeacherController) AddEvaluate(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	phone := c.PostForm("phone")
	if phone == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "phone is empty",
		})
		return
	}
	param["phone"] = phone
	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	appKey := c.PostForm("appkey")
	if appKey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[appKey]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	evaluate1 := c.PostForm("evaluate")
	if evaluate1 == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "evaluate is empty",
		})
		return
	}
	param["evaluate"] = evaluate1
	evaluate := false
	if evaluate1 == "0" {
		evaluate = true
	}
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	err = s.TeacherService.AddEvaluate(phone, evaluate)
	if err != nil {
		log.Error("add evaluate error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "add evaluate error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *TeacherController) GetTeacher(c *gin.Context) {
	param := map[string]interface{}{}
	signal := c.PostForm("signal")
	if signal == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is empty",
		})
		return
	}
	timestamp := c.PostForm("timestamp")
	times, err := strconv.Atoi(timestamp)
	if err != nil {
		log.Errorf("change time error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "change time error",
		})
		return
	}
	if time.Now().Unix() > int64(times) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is old",
		})
		return
	}
	if timestamp == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "timestamp is empty",
		})
		return
	}
	param["timestamp"] = timestamp
	appKey := c.PostForm("appkey")
	if appKey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is empty",
		})
		return
	}
	appkey := s.Config.AppKeyMap[appKey]
	if appkey == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "appkey is wrong",
		})
		return
	}
	newGrade := c.PostForm("newGrade")
	if newGrade == ""{
		c.JSON(201, gin.H{
			"result":  1,
			"message": "newGrade is empty",
		})
		return
	}
	newSubject := c.PostForm("newSubject")
	if newSubject == ""{
		c.JSON(201, gin.H{
			"result":  1,
			"message": "newSubject is empty",
		})
		return
	}
	newMoney := c.PostForm("newMoney")
	if newMoney == ""{
		c.JSON(201, gin.H{
			"result":  1,
			"message": "newMoney is empty",
		})
		return
	}
	newSex := c.PostForm("newSex")
	if newSex == ""{
		c.JSON(201, gin.H{
			"result":  1,
			"message": "newSex is empty",
		})
		return
	}
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	resp,err := s.TeacherService.GetTeacher(newGrade, newSubject, newMoney, newSex)
	if err != nil {
		log.Error("add evaluate error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "add evaluate error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
		"data" :resp,
	})
	return
}