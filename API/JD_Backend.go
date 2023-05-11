package API

import (
	"JD_backend/API/def"
	"JD_backend/DAO"
	"JD_backend/DAO/mdDef"
	"JD_backend/Service"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// BatchPolishJobs
// @Tags 工作管理
// @Summary 批量擦亮工作
// @Description 对多个工作进行高亮标注
// @Param Authorization header string true "Bearer 用户令牌"
// @Param Ids body []string true "希望高亮的JobId组"
// @Router /api/BatchPolishJobs [GET]
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} def.ResponseForm
func BatchPolishJobs(ctx *gin.Context) {

}

// CollectJob
// @Tags 工作管理
// @Summary 收藏工作
// @Description 对某一工作进入收藏夹
// @Param Authorization header string true "Bearer 用户令牌"
// @Param JobId query string true "将要收藏的工作对应的JobId"
// @Param UserId query string true "UserId对应用户的收藏夹"
// @Router /api/CollectJob [POST]
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
func CollectJob(ctx *gin.Context) {

}

func DeCollectJob(ctx *gin.Context) {

}

// BatchPullOffJobs
// @Tags 工作管理
// @Summary 批量下架工作
// @Description 将选中的工作进行下架
// @Param Authorization header string true "Bearer 用户令牌"
// @Param JobIds query []string true "将要下架的JobId组"
// @Router /api/BatchPullOffJobs [GET]
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
func BatchPullOffJobs(ctx *gin.Context) {

}

// ListUploadedJobs
// @Tags 工作管理
// @Summary 列出用户已上传的工作
// @Description 列出用户已上传的工作
// @Param Authorization header string true "Bearer 用户令牌"
// @Param UserId query string true "将要列出工作的用户UserId"
// @Router /api/ListUploadedJobs [GET]
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
func ListUploadedJobs(ctx *gin.Context) {

}

// ListCollectedJobs
// @Tags 工作管理
// @Summary 列出用户收藏的工作
// @Description 列出用户收藏的工作
// @Param Authorization header string true "Bearer 用户令牌"
// @Param UserId query string true "将要列出收藏工作的用户UserId"
// @Router /api/ListCollectedJobs [GET]
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
func ListCollectedJobs(ctx *gin.Context) {

}

// ListViewedJobs
// @Tags 工作管理
// @Summary 列出最近浏览的工作
// @Description 列出最近浏览的工作
// @Param Authorization header string true "Bearer 用户令牌"
// @Param UserId query string true "将要列出最近浏览的用户UserId"
// @Router /api/ListViewedJobs [GET]
// @Produce json
// @Success 200 {object} ResponseForm
// @Failure 400 {object} string
func ListViewedJobs(ctx *gin.Context) {

}

// Login
// @tags 用户管理
// @Summary 用户登录验证
// @Description 用户登录使用JWT验证
// @Param name query string true "用户ID"
// @Param password query string true "用户密码"
// @Router /Login [GET]
// @Produce json
// @Success 200 {object} ResponseForm
// @Failure 400 {object} string
func Login(ctx *gin.Context) {
	var u def.UserReq
	if err := ctx.ShouldBindJSON(&u); err != nil {
		log.Println("login info error msg is " + err.Error())
		return
	}
	user, err := DAO.GetUserByName(u.Name)
	if err != nil {
		log.Println("dont get user from username " + err.Error())
		return
	}
	if err := user.Compare(u.PassWord); err != nil {
		log.Println("password error " + err.Error())
		return
	}

	token, err := Service.Sign(user.UserId, user.Name)

	if err != nil {
		log.Println("get token error " + err.Error())
	}

	userId2Token := &mdDef.TokenBasic{}
	result := DAO.MysqlDB.Where("id = ?", user.UserId).First(userId2Token)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 数据不存在，插入数据
		newToken := mdDef.TokenBasic{
			UserId: user.UserId,
			Token:  token,
		}
		DAO.MysqlDB.Create(&newToken)

	} else {
		userId2Token.Token = token
		DAO.MysqlDB.Save(&userId2Token)
	}

	ctx.JSON(http.StatusOK, def.ResponseForm{Code: "200", Msg: "login success", Data: token})
}
