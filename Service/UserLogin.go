package Service

import (
	"JD_backend/DAO"
	"JD_backend/DAO/mdDef"
	"JD_backend/Service/msDef"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func Login(name string, password string) (interface{}, error) {
	token, err := DAO.Login(name, password)
	if err != nil {
		log.Println("login error in Service layer " + err.Error())
		return nil, err
	}
	return token, nil
}

func WXLogin(code string) (interface{}, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	var appId = ""
	var secret = ""
	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, appId, secret, code)

	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	wxResp := msDef.WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	if wxResp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%d  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}

	token, err := DAO.WXLogin(wxResp.OpenId, wxResp.SessionKey)
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
