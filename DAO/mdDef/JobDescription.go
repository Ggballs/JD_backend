package mdDef

import (
	"gorm.io/gorm"
	"time"
)

// JobDescription 职位描述
type JobDescription struct {
	gorm.Model
	JobId              string `json:"job_id" gorm:"not null; index"`
	InternTimeInMonths int    `json:"intern_time_in_months"`
	WorkDay            int
	BasePosition       string
	Degree             string
	JobType            string
	Industry           string
	JobName            string
	CompanyName        string
	//具体地点和base的区别？
	UploadUserId   string `json:"upload_user_id"`
	isShow         bool   `json:"is_show" gorm:"column:is_show"`
	PolishedTime   time.Time
	CollectedTimes int
}
