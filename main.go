package main

import (
	"JD_backend/V0"
	_ "JD_backend/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

var BasicRouter *gin.Engine

func init() {
	BasicRouter = gin.Default()
	V0.Register(BasicRouter)
	// test push
}
func main() {
	BasicRouter.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	BasicRouter.Run(":8080")
}
