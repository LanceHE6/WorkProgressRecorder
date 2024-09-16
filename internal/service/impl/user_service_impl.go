package impl

import (
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
)

// UserServiceImpl
//
//	@Description: 用户服务实现
type UserServiceImpl struct {
}

// Login
//
//	@Description: 登录
//	@receiver s UserServiceImpl 服务
//	@param account 账号
//	@param password 密码
//	@return *pkg.Response 返回结果
func (s UserServiceImpl) Login(account, password string) *pkg.Response {
	userRepo := repo.NewUserRepository()
	user := userRepo.SelectByAccountAndPsw(account, password)
	if user == nil {
		return pkg.FailedResponse(1, "账号或密码错误")
	}
	sessionID := pkg.GenerateUUID(8)
	token, err := pkg.GenerateToken(user.ID, user.Permission, user.Name, sessionID)
	if err != nil {
		return pkg.ErrorResponse(-1, "生成token失败", err)
	}
	// 更新session_id
	err = userRepo.UpdateSessionID(user.ID, sessionID)
	if err != nil {
		return pkg.ErrorResponse(-2, "更新session_id失败", err)
	}
	return pkg.SuccessResponse(map[string]any{
		"token": token,
		"user":  user,
	})
}

// Import
//
//	@Description: 导入用户
//	@receiver s UserServiceImpl
//	@param users 用户列表
//	@return *pkg.Response 返回结果
func (s UserServiceImpl) Import(users []model.User) *pkg.Response {
	userRepo := repo.NewUserRepository()
	errs := make(map[string]string)
	for _, user := range users {
		if len(user.Account) != 10 {
			errs[user.Account] = "账号长度必须为10位学号"
			continue
		}
		err := userRepo.Insert(user)
		if err != nil {
			errs[user.Account] = err.Error()
		}
	}
	if len(errs) > 0 {
		return pkg.NewResponse(1, "导入存在错误", map[string]any{
			"errs": errs,
		})
	} else {
		return pkg.SuccessResponse(nil)
	}
}

// UpdatePassword
//
//	@Description: 修改密码
//	@receiver s UserServiceImpl
//	@param id 用户id
//	@param oldPsw 原密码
//	@param newPsw 新密码
//	@return *pkg.Response 返回结果
func (s UserServiceImpl) UpdatePassword(id int64, oldPsw, newPsw string) *pkg.Response {
	user := repo.NewUserRepository().SelectByID(id)
	if user == nil {
		return pkg.FailedResponse(-1, "用户不存在")
	}
	if !pkg.CheckPsw(user.Password, oldPsw) {
		return pkg.FailedResponse(-2, "原密码错误")
	}
	err := repo.NewUserRepository().UpdatePassword(id, newPsw)
	if err != nil {
		return pkg.ErrorResponse(-3, "修改密码失败", err)
	} else {
		return pkg.SuccessResponse(nil)
	}
}
