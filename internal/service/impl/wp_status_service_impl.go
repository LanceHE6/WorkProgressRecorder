package impl

import (
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
)

type WPStatusServiceImpl struct {
}

// GetStatus
//
//	@Description: 获取就业状态列表
//	@receiver r *WPStatusServiceImpl
//	@return *pkg.Response 返回信息
func (r *WPStatusServiceImpl) GetStatus() *pkg.Response {
	status, err := repo.NewWpStatusRepo().ListStatus()
	if err != nil {
		return pkg.ErrorResponse(1, "获取状态失败", err)
	}
	return pkg.SuccessResponse(map[string]any{
		"status": status,
	})
}
