package DAO

import (
	"JD_backend/DAO/mdDef"
	"errors"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/utils"
	"log"
	"time"
)

func CollectJob(userId, jobId string) error {
	tx := MysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	JobDetail := mdDef.JobDescription{}
	r := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("job_id", jobId).Find(&JobDetail)
	if r.Error != nil {
		log.Println("DB transaction locking error : " + r.Error.Error())
		return r.Error
	}

	r = tx.Where("job_id", jobId).Find(&mdDef.JobDescription{})
	if r.Error != nil {
		log.Println("DB transaction finding error : " + r.Error.Error())
		return r.Error
	}

	r = tx.Model(&mdDef.JobDescription{}).Where("job_id", jobId).Update("collected_time", JobDetail.CollectedTimes+1)
	if r.Error != nil {
		log.Println("DB transaction update error : " + r.Error.Error())
		return r.Error
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("mysql transaction error : " + err.Error())
		return err
	}

	var collectionDetail mdDef.Collection
	err := MysqlDB.Where("user_id = ? AND job_id = ?", userId, jobId).Find(&collectionDetail).Error
	if err != nil {
		log.Println("CollectJob error in Finding collectionDetail in DAO Layer: " + err.Error())
		return err
	}
	if collectionDetail.JobId == jobId {
		return nil
	}

	err = MysqlDB.Create(&mdDef.Collection{UserId: userId, JobId: jobId}).Error
	if err != nil {
		log.Println("CollectJob error in Creating collectionDetail in DAO Layer: " + err.Error())
		return err
	}
	return nil
}

func DeCollectJob(userId, jobId string) error {
	tx := MysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	JobDetail := mdDef.JobDescription{}
	r := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("job_id", jobId).Find(&JobDetail)
	if r.Error != nil {
		log.Println("DeCollectJob : DB transaction locking error : " + r.Error.Error())
		return r.Error
	}

	r = tx.Where("job_id", jobId).Find(&mdDef.JobDescription{})
	if r.Error != nil {
		log.Println("DeCollectJob : DB transaction finding error : " + r.Error.Error())
		return r.Error
	}

	r = tx.Model(&mdDef.JobDescription{}).Where("job_id", jobId).Update("collected_time", JobDetail.CollectedTimes-1)
	if r.Error != nil {
		log.Println("DeCollectJob : DB transaction update error : " + r.Error.Error())
		return r.Error
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("DeCollectJob : mysql transaction error : " + err.Error())
		return err
	}

	var collectionDetail mdDef.Collection
	err := MysqlDB.Where("user_id = ? AND job_id = ?", userId, jobId).Find(&collectionDetail).Error
	if err != nil {
		log.Println("DeCollectJob error in Finding collectionDetail in DAO Layer: " + err.Error())
		return err
	}
	if collectionDetail.JobId != jobId {
		log.Println("DeCollectJob error in DAO Layer: " + "userId " + userId + " has no collection history on jobId " + jobId)
	}

	err = MysqlDB.Where("user_id = ? AND job_id = ?", userId, jobId).Delete(&mdDef.Collection{}).Error
	if err != nil {
		log.Println("DeCollectJob error in Delete collection in DAO Layer: " + err.Error())
		return err
	}
	return nil
}

func BatchPolishJobs(userId string, jobIds []string) error {
	ok, err := CheckJobsAccess(userId, jobIds)
	if err != nil {
		log.Println("BatchPolishJobs error in DAO Layer: " + err.Error())
		return err
	}
	if ok == false {
		return errors.New("BatchPullOffJobs error in DAO Layer: " + "userId " + userId + " has no access to jobIds")
	}
	for _, jobId := range jobIds {
		err = PolishJob(jobId)
		if err != nil {
			log.Println("BatchPolishJobs error in DAO Layer: " + err.Error())
			return err
		}
	}
	return nil
}

func PolishJob(jobId string) error {
	err := MysqlDB.Model(&mdDef.JobDescription{}).Where("job_id = ?", jobId).Update("polished_time", time.Now()).Error
	if err != nil {
		log.Println("PolishJob error in DAO Layer: " + err.Error())
		return err
	}
	return nil
}

func BatchPullOffJobs(userId string, jobIds []string) error {
	ok, err := CheckJobsAccess(userId, jobIds)
	if err != nil {
		log.Println("BatchPullOffJobs error in DAO Layer: " + err.Error())
		return err
	}
	if ok == false {
		return errors.New("BatchPullOffJobs error in DAO Layer: " + "userId " + userId + " has no access to jobIds")
	}
	for _, jobId := range jobIds {
		err = PullOffJob(jobId)
		if err != nil {
			log.Println("BatchPullOffJobs error in DAO Layer: " + err.Error())
			return err
		}
	}
	return nil
}
func PullOffJob(jobId string) error {
	err := MysqlDB.Model(&mdDef.JobDescription{}).Where("job_id = ?", jobId).Update("is_show", false).Error
	if err != nil {
		log.Println("PullOffJob error in DAO Layer: " + err.Error())
		return err
	}
	return nil
}

func CheckJobsAccess(userId string, jobIds []string) (bool, error) {
	uploadedJobsId, err := ListUploadedJobIds(userId)
	if err != nil {
		log.Println("CollectJob error in DAO Layer: " + err.Error())
		return false, err
	}
	for _, jobId := range jobIds {
		if utils.Contains(uploadedJobsId, jobId) == false {
			return false, nil
		}
	}
	return true, nil
}

func ListUploadedJobIds(userId string) ([]string, error) {
	var jobs []*mdDef.JobDescription
	err := MysqlDB.Where("upload_user_id = ?", userId).Find(&jobs).Error
	if err != nil {
		log.Println("ListUploadedJobs error in DAO Layer: " + err.Error())
		return nil, err
	}
	var jobIds []string
	for _, job := range jobs {
		jobIds = append(jobIds, job.JobId)
	}
	return jobIds, nil
}
