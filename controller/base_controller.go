package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/service/config"
	"strconv"
)

//os=%@; osver=%@; model=%@; appver=%@; appcode=%@;
type BaseController struct {
	Config *config.AppConfig
}

func (s *BaseController) GetOs(c *gin.Context) (clientIos int8, clientAndroid int8, clientWeb int8) {
	clientIos = -1
	clientAndroid = -1
	clientWeb = -1
	os := c.GetHeader("Os")
	if os == "ios" {
		clientIos = 0
	} else if os == "android" {
		clientAndroid = 0
	} else if os == "web" {
		clientWeb = 0
	}
	return clientIos, clientAndroid, clientWeb
}

//系统版本 10.3
func (s *BaseController) GetOsVer(c *gin.Context) string {
	return c.GetHeader("Osver")
}

//版本　 iphone4, iphone5
func (s *BaseController) GetModel(c *gin.Context) string {
	return c.GetHeader("Model")
}

//app的版本1.0.0
func (s *BaseController) GetAppVer(c *gin.Context) string {
	return c.GetHeader("Appver")
}

//app identifier
func (s *BaseController) GetIdentifier(c *gin.Context) string {
	return c.GetHeader("Identifier")
}

//app的code 1120
func (s *BaseController) GetAppCode(c *gin.Context) string {
	return c.GetHeader("Appcode")
}

func (s *BaseController) ParamInt64(c *gin.Context, key string) (int64, error) {
	paramValue := c.Param(key)
	if paramValue == "" {
		return -1, errors.New(key + " is empty")
	}
	return strconv.ParseInt(paramValue, 10, 64)
}

func (s *BaseController) DefaultQueryInt64(c *gin.Context, key string) (int64, error) {
	paramValue := c.DefaultQuery(key, "")
	if paramValue == "" {
		return -1, errors.New(key + " is empty")
	}
	return strconv.ParseInt(paramValue, 10, 64)
}

func (s *BaseController) DefaultQueryInt641(c *gin.Context, key string, value string) (int64, error) {
	paramValue := c.DefaultQuery(key, value)
	if paramValue == "" {
		return -1, errors.New(key + " is empty")
	}
	return strconv.ParseInt(paramValue, 10, 64)
}

func (s *BaseController) PostFormInt64(c *gin.Context, key string) (int64, error) {
	param := c.PostForm(key)
	if param == "" {
		return -1, errors.New(key + " is empty")
	}
	ret, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return -1, err
	}
	return ret, nil
}

func (s *BaseController) DefaultPostFormInt64(c *gin.Context, key string, value int64) (int64, error) {
	param := c.PostForm(key)
	if param == "" {
		return value, nil
	}
	ret, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return -1, err
	}
	return ret, nil
}

func (s *BaseController) DefaultPostFormFloat64(c *gin.Context, key string, value float64) (float64, error) {
	param := c.PostForm(key)
	if param == "" {
		return value, nil
	}
	ret, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return -1, err
	}
	return ret, nil
}
