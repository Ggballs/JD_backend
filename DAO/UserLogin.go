package DAO

import (
	"JD_backend/DAO/mdDef"
	"JD_backend/Service/msDef"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/gbrlsnchs/jwt/v3"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

var privateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
var publicKey = &privateKey.PublicKey
var hs = jwt.NewES256(jwt.ECDSAPublicKey(publicKey), jwt.ECDSAPrivateKey(privateKey))

func Sign(id string, auxiliary string) (string, error) {
	now := time.Now()
	pl := msDef.LoginToken{
		Payload: jwt.Payload{
			Issuer:         "lian_xu",
			Subject:        "login",
			Audience:       jwt.Audience{},
			ExpirationTime: jwt.NumericDate(now.Add(7 * 24 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Second)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          uuid.NewV4().String(),
		},
		ID:        id,
		Auxiliary: auxiliary,
	}
	token, err := jwt.Sign(pl, hs)
	return string(token), err
}

func Verify(token []byte) (*msDef.LoginToken, error) {
	pl := &msDef.LoginToken{}
	_, err := jwt.Verify(token, hs, pl)
	return pl, err
}

func GetUserByName(userName string) (*mdDef.UserBasic, error) {
	var user *mdDef.UserBasic
	result := MysqlDB.Where("name = ?", userName).First(&user)
	if result.Error != nil {
		log.Println("DB search error + ", result.Error.Error())
		return nil, result.Error
	}
	return user, nil
}

func Login(name string, password string) (interface{}, error) {
	user, err := GetUserByName(name)
	if err != nil {
		log.Println("dont get user from username " + err.Error())
		return nil, err
	}
	if err := user.Compare(password); err != nil {
		log.Println("password error " + err.Error())
	}

	token, err := Sign(user.UserId, user.Name)
	if err != nil {
		log.Println("get token error " + err.Error())
	}

	userId2token := mdDef.TokenBasic{}
	userId2token.UserId = user.UserId
	userId2token.Token = token
	MysqlDB.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"token"}),
	}).Create(&userId2token)

	return token, err
}

func WXLogin(userId string, sessionKey string) (interface{}, error) {
	token, err := Sign(userId, sessionKey)
	if err != nil {
		log.Println("get token error " + err.Error())
	}

	userId2token := mdDef.TokenBasic{}
	userId2token.UserId = userId
	userId2token.Token = token
	MysqlDB.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"token"}),
	}).Create(&userId2token)

	return token, err
}

func GetUserInfoByToken(token string) (*mdDef.UserBasic, error) {
	var token2user *mdDef.TokenBasic
	err := MysqlDB.Where("token = ?", token).First(&token2user).Error
	if err != nil {
		log.Println("get userInfoByHeader error in DAO layer " + err.Error())
		return nil, err
	}

	var user *mdDef.UserBasic
	err = MysqlDB.Where("id = ?", token2user.UserId).First(&user).Error
	if err != nil {
		log.Println("get userInfoByHeader error in DAO layer " + err.Error())
		return nil, err
	}
	return user, nil
}
