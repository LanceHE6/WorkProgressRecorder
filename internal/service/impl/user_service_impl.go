package impl

import (
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
	"fmt"
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
	fmt.Println("service:", account, password)
	user := userRepo.SelectByAccountAndPsw(account, password)
	fmt.Println(s)
	if user == nil {
		return pkg.SuccessResponse("账号或密码错误")
	}
	sessionID := pkg.GenerateUUID(8)
	token, err := pkg.GenerateToken(user.ID, user.Permission, user.Name, sessionID)
	if err != nil {
		return pkg.ErrorResponse(-1, "生成token失败")
	}
	return pkg.SuccessResponse(map[string]any{
		"token": token,
		"user":  user,
	})
}
