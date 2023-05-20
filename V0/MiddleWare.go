package V0

import (
	"JD_backend/API/def"
	"JD_backend/Service"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		headerList := strings.Split(header, " ")
		if len(headerList) != 2 {
			err := errors.New("unable to parse Auto")
			log.Println("header error " + err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, def.ResponseForm{Code: http.StatusUnauthorized, Msg: "unable to parse auto"})
			return
		}
		t := headerList[0]
		content := headerList[1]
		if t != "Bearer" {
			err := errors.New("error auto type")
			log.Println("header error " + err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, def.ResponseForm{Code: http.StatusUnauthorized, Msg: "error auto type"})
			return
		}

		if _, err := Service.Verify([]byte(content)); err != nil {
			err := errors.New("error token")
			log.Println("header error " + err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, def.ResponseForm{Code: http.StatusUnauthorized, Msg: "error token"})
			return
		}
		ctx.Set("AutoToken", content)
		ctx.Next()
	}
}
