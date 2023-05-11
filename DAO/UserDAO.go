package DAO

import (
	"JD_backend/DAO/mdDef"
	"log"
)

func GetUserByName(userName string) (*mdDef.UserBasic, error) {
	user := &mdDef.UserBasic{}
	result := MysqlDB.Where("name = ?", userName).First(user)
	if result.Error != nil {
		log.Println("DB search error + ", result.Error.Error())
		return nil, result.Error
	}
	return user, nil
}
