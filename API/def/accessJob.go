package def

type CollectJobRequest struct {
	UserId string
	JobId  string
}

type BatchPolishJobsRequest struct {
	UserId string
	JobIds []string
}

type BatchPullOffJobs struct {
	UserId string
	JobIds []string
}
