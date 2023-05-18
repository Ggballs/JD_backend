package Service

import (
	"JD_backend/DAO"
	"JD_backend/DAO/mdDef"
	"encoding/json"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	DAO.DBinit()
	viewedJobs := []mdDef.ViewedJob{
		{
			"1",
			"1",
		},
		{
			"2",
			"2",
		},
	}
	collectedJobs := []mdDef.CollectedJob{
		{
			"1",
			"1",
		},
		{
			"2",
			"2",
		},
	}
	jsonViewedJobs, err := json.Marshal(viewedJobs)
	jsonCollectedJobs, err := json.Marshal(collectedJobs)
	users := []mdDef.UserBasic{
		{
			UserId:        "1",
			Name:          "aliyah",
			PassWord:      "423319",
			ViewedJobs:    jsonViewedJobs,
			CollectedJobs: jsonCollectedJobs,
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
	err = DAO.MysqlDB.Create(&users).Error
	if err != nil {
		log.Println("add to database error " + err.Error())
		return
	}
}
