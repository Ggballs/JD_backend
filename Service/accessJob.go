package Service

import (
	"JD_backend/DAO"
	"JD_backend/DAO/mdDef"
	"log"
)

func CollectJob(userId, jobId string) error {
	err := DAO.CollectJob(userId, jobId)
	if err != nil {
		log.Println("CollectJob error : " + err.Error())
		return err
	}
	return nil
}

func DeCollectJob(userId, jobId string) error {
	err := DAO.DeCollectJob(userId, jobId)
	if err != nil {
		log.Println("DeCollectJob error : " + err.Error())
		return err
	}
	return nil
}

func BatchPolishJobs(userId string, jobIds []string) error {
	err := DAO.BatchPolishJobs(userId, jobIds)
	if err != nil {
		log.Println("BatchPolishJobs error : " + err.Error())
		return err
	}
	return nil
}

func BatchPullOffJobs(userId string, jobIds []string) error {
	err := DAO.BatchPullOffJobs(userId, jobIds)
	if err != nil {
		log.Println("BatchPullOffJobs error : " + err.Error())
		return err
	}
	return nil
}

func ListViewedJobs(userId string) ([]mdDef.JobDescription, error) {
	Jobs, err := DAO.ListViewedJobs(userId)
	if err != nil {
		log.Println("ListViewJobs error : " + err.Error())
		return nil, err
	}
	return Jobs, nil
}

func ListUploadedJobs(userId string) ([]mdDef.JobDescription, error) {
	jobs, err := DAO.ListUploadedJobs(userId)
	if err != nil {
		log.Println("ListUploadedJobs error ：" + err.Error())
		return nil, err
	}
	return jobs, nil
}

func ListCollectedJobs(userId string) ([]mdDef.JobDescription, error) {
	jobs, err := DAO.ListCollectedJobs(userId)
	if err != nil {
		log.Println("ListUploadedJobs error ：" + err.Error())
		return nil, err
	}
	return jobs, nil
}
