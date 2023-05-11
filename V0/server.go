package V0

import (
	"JD_backend/API"
	"github.com/gin-gonic/gin"
)

func Register(BasicRouter *gin.Engine) {
	BasicRouter.GET("login", API.Login)
	apiGroup := BasicRouter.Group("api")
	jobsGroup := apiGroup.Group("jobs")
	jobsGroup.POST("/polish", API.BatchPolishJobs)
	jobsGroup.POST("/collect", API.CollectJob)
	jobsGroup.POST("/pull-off", API.BatchPullOffJobs)
	apiGroup.DELETE("/de-collect", API.DeCollectJob)
	apiGroup.GET("/ListViewedJobs", API.ListViewedJobs)
	apiGroup.GET("/ListUploadedJobs", API.ListUploadedJobs)
	apiGroup.GET("/ListCollectedJobs", API.ListCollectedJobs)

}
