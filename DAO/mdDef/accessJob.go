package mdDef

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	UserId string `gorm:"not null; index; column:user_id" json:"user_id"`
	JobId  string `gorm:"not null; index; column:job_id" json:"job_id"`
}
