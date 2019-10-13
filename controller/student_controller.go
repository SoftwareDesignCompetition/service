package controller

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/service/middleware"
	"github.com/service/model"
	"github.com/service/service"
	"strconv"
	"time"
)

type StudentController struct {
	BaseController
	StudentService *service.StudentService
}

func NewStudentController(context *service.AppContext) (*StudentController, error) {
	studentService := context.Services["studentService"].(*service.StudentService)
	return &StudentController{
		BaseController: BaseController{
			Config: context.Config,
		},
		StudentService: studentService,
	}, nil
}
func (s *StudentController) Register(c *gin.Context) {
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
	name := c.PostForm("name")
	if name == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "name is empty",
		})
		return
	}
	param["name"] = name
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
			"message": "grand is empty",
		})
		return
	}
	param["grade"] = grade
	subject := c.PostForm("subject")
	if subject == "" {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "subject is empty",
		})
		return
	}
	param["subject"] = subject
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
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	student := model.Student{
		Phone:   phone,
		Name:    name,
		Subject: subject,
		Address: address,
		Grade:   grade,
	}
	resp, err := s.StudentService.GetAllPhone() //检验是否已经注册
	for _, i := range resp {
		if i == student.Phone {
			c.JSON(201, gin.H{
				"result":  1,
				"message": "Illegal double registration", //非法双重注册
			})
			return
		}
	}
	err = s.StudentService.Register(student)
	if err != nil {
		log.Errorf("studentRegister error %v", err)
		c.JSON(400, gin.H{
			"result":  2,
			"message": "studentRegister error",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
	})
	return
}

func (s *StudentController) GetStudent(c *gin.Context) {
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
		log.Error("change time error")
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
	if !middleware.CheckSign(appkey, param, signal) {
		c.JSON(201, gin.H{
			"result":  1,
			"message": "signal is wrong",
		})
		return
	}
	data, err := s.StudentService.GetStudentFormPhone(phone)
	if err != nil {
		log.Error("search student error")
		c.JSON(400, gin.H{
			"result":  2,
			"message": "search student error",
		})
		return
	}
	if data.Phone == "" {
		c.JSON(200, gin.H{
			"result":  0,
			"message": "success",
			"data":    "",
		})
		return
	}
	c.JSON(200, gin.H{
		"result":  0,
		"message": "success",
		"data":    fmt.Sprint(data),
	})
	return

}

func (s *StudentController) ChangePhone(c *gin.Context) {
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
		log.Error("change time error")
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
	err = s.StudentService.ChangePhone(phone, new_phone)
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

func (s *StudentController) ChangeAddress(c *gin.Context) {
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
		log.Error("change time error")
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
	err = s.StudentService.ChangeAddress(phone, address)
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

func (s *StudentController) ChangeGrade(c *gin.Context) {
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
		log.Error("change time error")
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
	err = s.StudentService.ChangeGrade(phone, grade)
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

func (s *StudentController) ChangeSubject(c *gin.Context) {
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
		log.Error("change time error")
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
	err = s.StudentService.ChangeSubject(phone, subject)
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

func (s *StudentController) ChangeName(c *gin.Context) {
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
		log.Error("change time error")
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
	err = s.StudentService.ChangeName(phone, name)
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
