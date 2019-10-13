package middleware

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
	"time"

	log "github.com/cihub/seelog"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Info(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Info(status)
	}
}

func CheckSign(appKey string, params map[string]interface{}, sign string) bool {
	result, needSignStr := DoSign(appKey, params)
	if result == sign {
		return true
	} else {
		log.Warnf("sign string %v sign %v, result %v", needSignStr, sign, result)
	}
	return false
}

func DoSign(appKey string, params map[string]interface{}) (string, string) {
	var keys []string
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var needSignStr = "app_key=" + appKey
	for _, k := range keys {
		needSignStr += "&" + k + "=" + fmt.Sprintf("%v", params[k])
	}
	h := sha256.New()
	h.Write([]byte(needSignStr))
	return fmt.Sprintf("%x", h.Sum(nil)), needSignStr
}
