package mdDef

import (
	"gorm.io/gorm"
	"time"
)

// JobDescription 职位描述
type JobDescription struct {
	gorm.Model
	JobId              string    `json:"job_id" gorm:"not null; index"`
	InternTimeInMonths int       `json:"intern_time_in_months"`
	WorkDay            int       `json:"work_day"`
	BasePosition       string    `json:"base_position"`
	Degree             string    `json:"degree"`
	JobType            string    `json:"job_type"`
	Industry           string    `json:"industry"`
	JobName            string    `json:"job_name"`
	CompanyName        string    `json:"company_name"`
	UploadUserId       string    `json:"upload_user_id"`
	IsShow             bool      `json:"is_show" gorm:"column:is_show"`
	PolishedTime       time.Time `json:"polished_time"`
	CollectedTimes     int       `json:"collected_times"`
}
