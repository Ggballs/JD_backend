package V0

import (
	"JD_backend/mAPI"
	"github.com/gin-gonic/gin"
)

func Register(BasicRouter *gin.Engine) {
	BasicRouter.GET("/api/PolishJob", mAPI.PolishJob)
	BasicRouter.GET("/api/BatchCollectJobs", mAPI.BatchPolishJobs)
	BasicRouter.GET("/api/CollectJob", mAPI.CollectJob)
	BasicRouter.GET("/api/ListViewedJobs", mAPI.ListViewedJobs)
	BasicRouter.GET("/api/ListUploadedJobs", mAPI.ListUploadedJobs)
	BasicRouter.GET("/api/ListCollectedJobs", mAPI.ListCollectedJobs)
	BasicRouter.GET("/api/BatchPullOff", mAPI.BatchPullOff)
}
