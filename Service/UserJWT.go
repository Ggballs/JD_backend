package Service

import (
	"JD_backend/DAO"
	"JD_backend/DAO/mdDef"
	"JD_backend/Service/msDef"
	"log"
)

func Login(name string, password string) (interface{}, error) {
	token, err := DAO.Login(name, password)
	if err != nil {
		log.Println("login error in Service layer " + err.Error())
		return nil, err
	}
	return token, nil
}

func Verify(token []byte) (*msDef.LoginToken, error) {
	pl, err := DAO.Verify(token)
	if err != nil {
		log.Println("login error in Service layer " + err.Error())
		return nil, err
	}
	return pl, nil
}

func GetUserInfoByToken(token string) (*mdDef.UserBasic, error) {
	user, err := DAO.GetUserInfoByToken(token)
	if err != nil {
		log.Println("get userInfoByHeader error in Service layer " + err.Error())
		return nil, err
	}
	return user, err
}
