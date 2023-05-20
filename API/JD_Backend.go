package API

import (
	"JD_backend/API/def"
	"JD_backend/Service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// BatchPolishJobs
// @Tags 工作管理
// @Summary 批量擦亮工作
// @Description 对多个工作进行高亮标注
// @Accept application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param request body def.BatchPolishJobsRequest true "请求"
// @Router /api/jobs/polish [POST]
// @Produce json
// @Success 200 {object} def.ResponseForm
// @Failure 400 {object} def.ResponseForm
func BatchPolishJobs(ctx *gin.Context) {
	req := new(def.BatchPolishJobsRequest)
	if err := ctx.ShouldBind(req); err != nil {
		log.Println("BatchPolishJobs Error in API layer : " + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	token, ok := ctx.Get("AutoToken")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, def.ResponseForm{Code: http.StatusInternalServerError, Msg: "get token error"})
		return
	}
	userInfo, err := Service.GetUserInfoByToken(token.(string))

	if err != nil {
		log.Println("BatchPolishJobs Error in API layer : " + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	if err := Service.BatchPolishJobs(userInfo.UserId, req.JobIds); err != nil {
		log.Println("BatchPolishJobs Error in API layer : " + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, def.ResponseForm{
		Code: http.StatusOK,
		Msg:  "success",
	})
}

// CollectJob
// @Tags 工作管理
// @Summary 收藏工作
// @Description 对某一工作进入收藏夹
// @Param Authorization header string true "Bearer 用户令牌"
// @Param request body def.CollectJobRequest true "请求"
// @Router /api/jobs/collect [POST]
// @Produce json
// @Success 200 {object} def.ResponseForm
// @Failure 400 {object} def.ResponseForm
func CollectJob(ctx *gin.Context) {
	req := new(def.CollectJobRequest)
	if err := ctx.ShouldBind(req); err != nil {
		log.Println("CollectJob Error in API layer : " + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	token, ok := ctx.Get("AutoToken")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, def.ResponseForm{Code: http.StatusInternalServerError, Msg: "get token error"})
		return
	}
	userInfo, err := Service.GetUserInfoByToken(token.(string))
	if err != nil {
		log.Println("CollectJob Error in API layer :" + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	if err := Service.CollectJob(userInfo.UserId, req.JobId); err != nil {
		log.Println("CollectJob Error in API layer :" + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, def.ResponseForm{
		Code: http.StatusOK,
		Msg:  "success",
	})
}

// DeCollectJob
// @Tags 工作管理
// @Summary 取消收藏工作
// @Description 对某一工作取消收藏
// @Param Authorization header string true "Bearer 用户令牌"
// @Param request body def.CollectJobRequest true "请求"
// @Router /api/jobs/de-collect [POST]
// @Produce json
// @Success 200 {object} def.ResponseForm
// @Failure 400 {object} def.ResponseForm
func DeCollectJob(ctx *gin.Context) {
	req := new(def.CollectJobRequest)
	if err := ctx.ShouldBind(req); err != nil {
		log.Println("DeCollectJob Error in API layer : " + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	token, ok := ctx.Get("AutoToken")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, def.ResponseForm{Code: http.StatusInternalServerError, Msg: "get token error"})
		return
	}
	userInfo, err := Service.GetUserInfoByToken(token.(string))
	if err != nil {
		log.Println("DeCollectJob Error in API layer :" + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	if err := Service.DeCollectJob(userInfo.UserId, req.JobId); err != nil {
		log.Println("DeCollectJob Error in API layer :" + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, def.ResponseForm{
		Code: http.StatusOK,
		Msg:  "success",
	})
}

// BatchPullOffJobs
// @Tags 工作管理
// @Summary 批量下架工作
// @Description 将选中的工作进行下架
// @Param Authorization header string true "Bearer 用户令牌"
// @Param request body def.BatchPullOffJobsRequest true "请求"
// @Router /api/jobs/pull-off [POST]
// @Produce json
// @Success 200 {object} def.ResponseForm
// @Failure 400 {object} def.ResponseForm
func BatchPullOffJobs(ctx *gin.Context) {
	req := new(def.BatchPullOffJobsRequest)
	if err := ctx.ShouldBind(req); err != nil {
		log.Println("BatchPullOffJobs Error in API layer : " + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	token, ok := ctx.Get("AutoToken")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, def.ResponseForm{Code: http.StatusInternalServerError, Msg: "get token error"})
		return
	}
	userInfo, err := Service.GetUserInfoByToken(token.(string))
	if err != nil {
		log.Println("BatchPullOffJobs Error in API layer :" + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	if err := Service.BatchPullOffJobs(userInfo.UserId, req.JobIds); err != nil {
		log.Println("BatchPullOffJobs Error in API layer :" + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, def.ResponseForm{
		Code: http.StatusOK,
		Msg:  "success",
	})
}

// ListUploadedJobs
// @Tags 工作管理
// @Summary 列出用户已上传的工作
// @Description 列出用户已上传的工作
// @Param Authorization header string true "Bearer 用户令牌"
// @Router /api/List-uploads [GET]
// @Produce json
// @Success 200 {object} def.ResponseForm{data=mdDef.JobDescription} “工作详情”
// @Failure 400 {object} def.ResponseForm
func ListUploadedJobs(ctx *gin.Context) {
	token, ok := ctx.Get("AutoToken")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, def.ResponseForm{Code: http.StatusInternalServerError, Msg: "get token error"})
		return
	}
	userInfo, err := Service.GetUserInfoByToken(token.(string))
	jobs, err := Service.ListUploadedJobs(userInfo.UserId)
	if err != nil {
		log.Println("login viewed job Error in API layer :" + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, def.ResponseForm{Code: http.StatusOK, Msg: "success", Data: jobs})
}

// ListCollectedJobs
// @Tags 工作管理
// @Summary 列出用户收藏的工作
// @Description 列出用户收藏的工作
// @Param Authorization header string true "Bearer 用户令牌"
// @Router /api/jobs/list-collections [GET]
// @Produce json
// @Success 200 {object} def.ResponseForm{data=mdDef.JobDescription} “工作详情”
// @Failure 400 {object} def.ResponseForm
func ListCollectedJobs(ctx *gin.Context) {
	token, ok := ctx.Get("AutoToken")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, def.ResponseForm{Code: http.StatusInternalServerError, Msg: "get token error"})
		return
	}
	userInfo, err := Service.GetUserInfoByToken(token.(string))
	jobs, err := Service.ListCollectedJobs(userInfo.UserId)
	if err != nil {
		log.Println("login viewed job Error in API layer :" + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, def.ResponseForm{Code: http.StatusOK, Msg: "success", Data: jobs})
}

// ListViewedJobs
// @Tags 工作管理
// @Summary 列出最近浏览的工作
// @Description 列出最近浏览的工作
// @Param Authorization header string true "Bearer 用户令牌"
// @Router /api/jobs/list-views [GET]
// @Produce json
// @Success 200 {object} def.ResponseForm{data=mdDef.JobDescription} “工作详情”
// @Failure 400 {object} def.ResponseForm
func ListViewedJobs(ctx *gin.Context) {
	token, ok := ctx.Get("AutoToken")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, def.ResponseForm{Code: http.StatusInternalServerError, Msg: "get token error"})
		return
	}
	userInfo, err := Service.GetUserInfoByToken(token.(string))
	jobs, err := Service.ListViewedJobs(userInfo.UserId)
	if err != nil {
		log.Println("login viewed job Error in API layer : " + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, def.ResponseForm{Code: http.StatusOK, Msg: "success", Data: jobs})
}

// Login
// @tags 用户管理
// @Summary 用户登录验证
// @Description 用户登录使用JWT验证
// @Param request body def.UserRequest true "请求"
// @Router /Login [POST]
// @Produce json
// @Success 200 {object} def.ResponseForm{data=msDef.Token} "token"
// @Failure 400 {object} def.ResponseForm
func Login(ctx *gin.Context) {
	var r def.UserRequest
	if err := ctx.ShouldBindJSON(&r); err != nil {
		log.Println("login Error in API layer : " + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	token, err := Service.Login(r.Name, r.PassWord)
	if err != nil {
		log.Println("login Error in API layer : " + err.Error())
		ctx.JSON(http.StatusBadRequest, def.ResponseForm{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, def.ResponseForm{Code: http.StatusOK, Msg: "success", Data: token.(string)})
}
