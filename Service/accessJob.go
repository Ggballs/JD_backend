package Service

import (
	"JD_backend/DAO"
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

func ListViewedJobs(userId string) ([]string, error) {
	JobIds, err := DAO.ListViewedJobIds(userId)
	if err != nil {
		log.Println("ListViewJobs error : " + err.Error())
		return nil, err
	}
	return JobIds, nil
}

func ListUploadedJobs(userId string) ([]string, error) {
	jobIds, err := DAO.ListUploadedJobIds(userId)
	if err != nil {
		log.Println("ListUploadedJobs error ：" + err.Error())
		return nil, err
	}
	return jobIds, nil
}

func ListCollectedJobs(userId string) ([]string, error) {
	jobIds, err := DAO.ListCollectedJobIds(userId)
	if err != nil {
		log.Println("ListUploadedJobs error ：" + err.Error())
		return nil, err
	}
	return jobIds, nil
}
