package mdDef

// JobDescription 职位描述
type JobDescription struct {
	JobId              int
	InternTimeInMonths int
	WorkDay            int
	BasePosition       string
	Degree             string
	JobType            string
	Industry           string
	JobName            string
	CompanyName        string
	//具体地点和base的区别？
	isShow         string
	CollectedTimes int
}
