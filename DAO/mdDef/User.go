package mdDef

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId string
	Name   string
}
