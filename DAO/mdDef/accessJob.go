package mdDef

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	UserId string
	JobId  string
}
