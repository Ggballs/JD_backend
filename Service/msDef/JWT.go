package msDef

import "github.com/gbrlsnchs/jwt/v3"

type LoginToken struct {
	jwt.Payload
	ID       string `json:"id"`
	UserName string `json:"userName"`
}

type Token struct {
	Token string `json:"token"`
}
