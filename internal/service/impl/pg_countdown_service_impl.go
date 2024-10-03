package impl

import (
	"WorkProgressRecord/config"
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// PgCountdownServiceImpl
//
//	@Description: 考研倒计时服务实现
type PgCountdownServiceImpl struct {
}

// GetPgCountdown
//
//	@Description: 获取考研倒计时
//	@receiver PgCountdownServiceImpl
//	@return int 倒计时天数
//	@return error 错误信息
func (PgCountdownServiceImpl) GetPgCountdown(context *gin.Context) {
	// 解析考研的日期字符串
	t, err := time.Parse("2006-01-02", config.GetPgeeTime())
	if err != nil {
		context.JSON(http.StatusOK, pkg.ErrorResponse(1, "解析考研日期失败", err))
		return
	}
	// 获取当前时间
	now := time.Now()
	// 计算时间差
	d := t.Sub(now)
	// 将时间差转换为天数
	days := int(d.Hours() / 24)
	context.JSON(http.StatusOK, pkg.SuccessResponse(map[string]any{
		"days": days,
	}))
}
