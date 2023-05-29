package def

import "JD_backend/DAO/mdDef"

type ResponseForm struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ListJobResponse struct {
	Jobs   []mdDef.JobDescription `json:"jobs"`
	Length int                    `json:"length"`
}
