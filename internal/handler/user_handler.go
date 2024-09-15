package handler

import (
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
