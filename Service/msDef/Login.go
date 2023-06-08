package msDef

import (
	"github.com/gbrlsnchs/jwt/v3"
)

type LoginToken struct {
	jwt.Payload
	ID        string `json:"id"`
	Auxiliary string `json:"Auxiliary"`
}

type Token struct {
	Token string `json:"token"`
}

type WXLoginResp struct {
	OpenId     string `json:"open_id"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"union_id"`
	ErrCode    int    `json:"err_code"`
	ErrMsg     string `json:"err_msg"`
}
