package Service

import (
	"JD_backend/Service/msDef"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/gbrlsnchs/jwt/v3"
	uuid "github.com/satori/go.uuid"
	"time"
)

var privateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
var publicKey = &privateKey.PublicKey
var hs = jwt.NewES256(jwt.ECDSAPublicKey(publicKey), jwt.ECDSAPrivateKey(privateKey))

func Sign(id string, username string) (string, error) {
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
		ID:       id,
		UserName: username,
	}
	token, err := jwt.Sign(pl, hs)
	return string(token), err
}

func Verify(token []byte) (*msDef.LoginToken, error) {
	pl := &msDef.LoginToken{}
	_, err := jwt.Verify(token, hs, pl)
	return pl, err
}
