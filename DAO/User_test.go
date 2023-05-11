package DAO

import (
	"JD_backend/DAO/mdDef"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	DBinit()

	users := []mdDef.UserBasic{
		{
			Name:     "aliyah",
			PassWord: "423319",
		},
	}

	for idx, _ := range users {
		if err := users[idx].Validate(); err != nil {
			log.Println("create user service interface error : " + err.Error())
			return
		}
	}

	for idx, _ := range users {
		if err := users[idx].Encrypt(); err != nil {
			log.Println("encrypt interface error : " + err.Error())
			return
		}
	}
	err := MysqlDB.Create(&users).Error
	if err != nil {
		log.Println("add to database error " + err.Error())
		return
	}
}
