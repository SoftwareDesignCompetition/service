package rediskey

import (
	"fmt"
)

//这里统一放redis key

func GetStudentRedisKey() string {
	return fmt.Sprintf("student_key")
}
