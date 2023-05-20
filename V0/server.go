package V0

import (
	"JD_backend/API"
	"github.com/gin-gonic/gin"
)

func Register(BasicRouter *gin.Engine) {
	BasicRouter.POST("login", API.Login)
	apiGroup := BasicRouter.Group("api")
	jobsGroup := apiGroup.Group("jobs")
	jobsGroup.Use(AuthJWT())
	{
		jobsGroup.POST("/polish", API.BatchPolishJobs)
		jobsGroup.POST("/collect", API.CollectJob)
		jobsGroup.POST("/pull-off", API.BatchPullOffJobs)
		jobsGroup.DELETE("/de-collect", API.DeCollectJob)
		jobsGroup.GET("/list-views", API.ListViewedJobs)
		jobsGroup.GET("/list-uploads", API.ListUploadedJobs)
		jobsGroup.GET("/list-collections", API.ListCollectedJobs)
	}
}
