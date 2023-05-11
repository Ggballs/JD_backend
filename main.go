package main

import (
	"JD_backend/DAO"
	"JD_backend/V0"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

var BasicRouter *gin.Engine

func init() {
	DAO.DBinit()
	BasicRouter = gin.Default()
	V0.Register(BasicRouter)
}
func main() {
	BasicRouter.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	BasicRouter.Run(":8080")
}
