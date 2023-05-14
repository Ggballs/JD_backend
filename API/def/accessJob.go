package def

type CollectJobRequest struct {
	UserId string
	JobId  string
}

type BatchPolishJobsRequest struct {
	UserId string
	JobIds []string
}

type BatchPullOffJobsRequest struct {
	UserId string
	JobIds []string
}

type ListViewedJobsRequest struct {
	UserId string `json:"user_id"`
}
