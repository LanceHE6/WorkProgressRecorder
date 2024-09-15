package handler

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/service"
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserHandler
//
//	@Description: 用户处理器
type UserHandler struct {
}

// loginRequest
//
//	@Description: 登录请求参数结构体
type loginRequest struct {
	Account  string `json:"account" form:"account" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// Login
//
//	@Description: 登录
//	@receiver h *UserHandler
//	@param context *gin.Context
func (h *UserHandler) Login(context *gin.Context) {
	var data loginRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(100, err.Error()))
		return
	}
	userService := service.NewUserService()
	context.JSON(http.StatusOK, userService.Login(data.Account, data.Password))
}

// importRequest
//
//	@Description: 导入用户请求参数结构体
type importRequest struct {
	Account string `json:"account" form:"account" binding:"required"`
	Name    string `json:"name" form:"name" binding:"required"`
}
type importUsersRequest struct {
	List []importRequest `json:"list" form:"list" binding:"required"`
}

// ImportUser
//
//	@Description: 导入用户
//	@receiver h *UserHandler
//	@param context *gin.Context
func (h *UserHandler) ImportUser(context *gin.Context) {
	var data importUsersRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(100, err.Error()))
		return
	}

	userList := data.List
	var users []model.User
	for _, userInfo := range userList {
		// 循环构建User对象
		var user model.User
		user.Account = userInfo.Account
		user.Name = userInfo.Name
		users = append(users, user)
	}
	userService := service.NewUserService()
	context.JSON(http.StatusOK, userService.Import(users))
}
