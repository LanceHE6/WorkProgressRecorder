package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
)

// EmplGoalServiceImpl
//
//	@Description: 就业目标服务实现
type EmplGoalServiceImpl struct {
}

// InsertAndUpdate
//
//	@Description: 添加或更新就业目标
//	@receiver e EmplGoalServiceImpl
//	@param emplGoal model.EmploymentGoal
//	@return *pkg.Response 返回响应
func (e EmplGoalServiceImpl) InsertAndUpdate(emplGoal model.EmploymentGoal) *pkg.Response {
	emplGoalRepo := repo.NewEmplGoalRepo()
	oldEmplGoal, err := emplGoalRepo.SelectByUID(emplGoal.UID)
	// 不存在则添加
	if err != nil || oldEmplGoal.ID == 0 {
		err = emplGoalRepo.Insert(emplGoal)
		if err != nil {
			return pkg.ErrorResponse(1, "添加失败", err)
		}
	} else {
		// 存在则更新
		emplGoal.ID = oldEmplGoal.ID
		err = emplGoalRepo.Update(emplGoal)
		if err != nil {
			return pkg.ErrorResponse(1, "更新失败", err)
		}
	}

	return pkg.SuccessResponse(nil)
}
