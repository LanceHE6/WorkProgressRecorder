package handler

import (
	"WorkProgressRecord/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PgCountdownHandler struct {
}

// GetPgCountdown
//
//	@Description: 获取考研倒计时
//	@receiver PgCountdownHandler
//	@param context
func (PgCountdownHandler) GetPgCountdown(context *gin.Context) {
	context.JSON(http.StatusOK, service.NewPgCountdownService().GetPgCountdown())
}
