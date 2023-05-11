package V0

import (
	"JD_backend/Service"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		headerList := strings.Split(header, " ")
		if len(headerList) != 2 {
			err := errors.New("unable to parse Auto")
			log.Println("header error " + err.Error())
			ctx.Abort()
			return
		}
		t := headerList[0]
		content := headerList[1]
		if t != "Bearer" {
			err := errors.New("error auto type")
			log.Println("header error " + err.Error())
			ctx.Abort()
			return
		}

		if _, err := Service.Verify([]byte(content)); err != nil {
			err := errors.New("error token")
			log.Println("header error " + err.Error())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
