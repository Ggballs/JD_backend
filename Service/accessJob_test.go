package Service

import (
	"JD_backend/DAO"
	"JD_backend/DAO/mdDef"
	"fmt"
	"testing"
	"time"
)

func initJD() {
	jobInfo := mdDef.JobDescription{
		JobId:        "testId",
		JobName:      "testName",
		UploadUserId: "testUser",
		PolishedTime: time.Now(),
	}
	err := DAO.MysqlDB.Create(&jobInfo).Error
	if err != nil {
		println("failed to create :" + err.Error())
	}
}

func init() {
	DAO.DBinit()
}
func TestCollect(t *testing.T) {
	//initJD()
	err := CollectJob("1111", "testId")
	if err != nil {
		println("service err " + err.Error())
		panic(t)
	}
}

func TestDeCollect(t *testing.T) {
	err := DeCollectJob("1111", "testId")
	if err != nil {
		println("service err " + err.Error())
		panic(t)
	}
}

func TestPolish(t *testing.T) {
	//initJD()
	err := BatchPolishJobs("testUser", []string{"testId"})
	if err != nil {
		println("service err " + err.Error())
		panic(t)
	}
}

func TestPullOff(t *testing.T) {
	err := BatchPullOffJobs("testUser", []string{"testId"})
	if err != nil {
		println("service err " + err.Error())
		panic(t)
	}
}

func TestShowCollections(t *testing.T) {
	var collections []mdDef.Collection
	err := DAO.MysqlDB.Find(&collections).Error
	if err != nil {
		println("dao err " + err.Error())
		panic(t)
	}
	for _, collection := range collections {
		fmt.Printf("%+v \n", collection)
	}
}

func TestDelete(t *testing.T) {
	DAO.MysqlDB.Unscoped().Where("id > 0").Delete(&mdDef.Collection{})
	DAO.MysqlDB.Unscoped().Where("id > 0").Delete(&mdDef.JobDescription{})

}
