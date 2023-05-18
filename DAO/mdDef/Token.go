package mdDef

import "gorm.io/gorm"

type TokenBasic struct {
	gorm.Model
	UserId string `json:"user_id" gorm:"not null; unique; column:user_id"`
	Token  string `json:"token" gorm:"not null; index; column:token"`
}

func (TokenBasic) TableName() string {
	return "token_basics"
}
