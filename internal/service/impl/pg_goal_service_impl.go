package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
)

// PGGoalServiceImpl
//
//	@Description: 考研目标服务实现
type PGGoalServiceImpl struct {
}

// InsertAndUpdate
//
//	@Description: 添加或更新考研目标
//	@receiver e PGGoalServiceImpl
//	@param emplGoal model.PostgraduateGoal
//	@return *pkg.Response 返回响应
func (p PGGoalServiceImpl) InsertAndUpdate(pgGoal model.PostgraduateGoal) *pkg.Response {
	pgGoalRepo := repo.NewPGGoalRepo()
	oldPgGoal, err := pgGoalRepo.SelectByUID(pgGoal.UID)
	if err != nil || oldPgGoal.ID == 0 {
		err = pgGoalRepo.Insert(pgGoal)
		if err != nil {
			return pkg.ErrorResponse(1, "添加失败", err)
		}
	} else {
		pgGoal.ID = oldPgGoal.ID
		err = pgGoalRepo.Update(pgGoal)
		if err != nil {
			return pkg.ErrorResponse(1, "更新失败", err)
		}
	}
	return pkg.SuccessResponse(nil)
}
