package impl

import (
	"WorkProgressRecord/config"
	"WorkProgressRecord/pkg"
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
func (PgCountdownServiceImpl) GetPgCountdown() *pkg.Response {
	// 解析考研的日期字符串
	t, err := time.Parse("2006-01-02", config.ServerConfig.PGEE.PGEE_TIME)
	if err != nil {
		return pkg.ErrorResponse(1, "解析考研日期失败", err)
	}
	// 获取当前时间
	now := time.Now()
	// 计算时间差
	d := t.Sub(now)
	// 将时间差转换为天数
	days := int(d.Hours() / 24)
	return pkg.SuccessResponse(map[string]any{
		"days": days,
	})
}
