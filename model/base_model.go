package model

type Student struct {
	Phone   string
	Name    string
	Subject string
	Address string
	Grade   string
	Good    int
	Own     int
}

type Teacher struct {
	Phone   string //手机号
	Name    string //姓名
	Gender  string //性别
	School  string //教师学校
	Grade   string //教师年级
	Major   string //教师专业
	Salary  string //薪资
	Address string //住址
	Times   string //时间
	Subject string //教授科目
	TeGrade string //教授年级
	Status  int    //是否接单
	Good    int    //好评次数
	Own     int    //总的接单次数
}
