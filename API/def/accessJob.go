package def

type CollectJobRequest struct {
	JobId string
}

type BatchPolishJobsRequest struct {
	JobIds []string
}

type BatchPullOffJobsRequest struct {
	JobIds []string
}

type ListViewedJobsRequest struct {
	UserId string `json:"user_id"`
}

type ListCollectedJobsRequest struct {
	UserId string `json:"user_id"`
}

type ListUploadedJobsRequest struct {
	UserId string `json:"user_id"`
}
