package impl

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
func (s UserServiceImpl) Login(context *gin.Context) {
	// loginRequest
	//
	//	@Description: 登录请求参数结构体
	type loginRequest struct {
		Account  string `json:"account" form:"account" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
	var data loginRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}
	userRepo := repo.NewUserRepository()
	user := userRepo.SelectByAccountAndPsw(data.Account, data.Password)
	if user == nil {
		context.JSON(http.StatusOK, pkg.FailedResponse(1, "账号或密码错误"))
		return
	}
	sessionID := pkg.GenerateUUID(8)
	token, err := pkg.GenerateToken(user.ID, user.Permission, user.Name, sessionID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, pkg.FailedResponse(1, "生成token失败"))
		return
	}
	// 更新session_id
	err = userRepo.UpdateSessionID(user.ID, sessionID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, pkg.FailedResponse(1, "更新session_id失败"))
		return
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(map[string]any{
		"token": token,
		"user":  user,
	}))
	return
}

// Import
//
//	@Description: 导入用户
//	@receiver s UserServiceImpl
//	@param users 用户列表
//	@return *pkg.Response 返回结果
func (s UserServiceImpl) Import(context *gin.Context) {
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
	var data importUsersRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
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
		context.JSON(http.StatusOK, pkg.NewResponse(1, "导入存在错误", map[string]any{
			"errs": errs,
		}))
		return

	} else {
		context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
		return
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
func (s UserServiceImpl) UpdatePassword(context *gin.Context) {
	// updatePasswordRequest
	// @Description: 修改密码请求参数结构体
	type updatePasswordRequest struct {
		OldPassword string `json:"old_password" form:"old_password" binding:"required"`
		NewPassword string `json:"new_password" form:"new_password" binding:"required"`
	}
	var data updatePasswordRequest
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}
	claims, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(-1, err.Error()))
	}
	user := repo.NewUserRepository().SelectByID(claims.ID)
	if user == nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(-1, "用户不存在"))
		return
	}
	if !pkg.CheckPsw(user.Password, data.OldPassword) {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(-2, "原密码错误"))
		return
	}
	err = repo.NewUserRepository().UpdatePassword(claims.ID, data.NewPassword)
	if err != nil {
		context.JSON(http.StatusInternalServerError, pkg.ErrorResponse(-3, "修改密码失败", err))
		return
	} else {
		context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
		return
	}
}

// GetUserTargetInfo
//
//	@Description: 获取用户目标信息
//	@receiver s UserServiceImpl
//	@param id 用户id
//	@return *pkg.Response 返回结果
func (s UserServiceImpl) GetUserTargetInfo(context *gin.Context) {
	userInfo, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-1, "无法获取用户id", err))
	}
	user := repo.NewUserRepository().SelectByID(userInfo.ID)
	if user == nil {
		context.JSON(http.StatusOK, pkg.FailedResponse(-1, "用户不存在"))
		return
	}
	if user.Direction == 0 {
		context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
		return
	}
	// 考研为目标
	if user.Direction == 1 {
		pgGoal, err := repo.NewPGGoalRepo().SelectByUID(userInfo.ID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, pkg.ErrorResponse(-2, "查询考研目标失败", err))
			return
		}
		context.JSON(http.StatusOK, pkg.SuccessResponse(map[string]any{
			"direction": user.Direction,
			"goal":      pgGoal,
		}))
		return
	}
	// 就业为目标
	if user.Direction == 2 {
		emplGoal, err := repo.NewEmplGoalRepo().SelectByUID(userInfo.ID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, pkg.ErrorResponse(-2, "查询就业目标失败", err))
			return
		}
		context.JSON(http.StatusOK, pkg.SuccessResponse(map[string]any{
			"direction": user.Direction,
			"goal":      emplGoal,
		}))
		return
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
}

// GetUserInfo
//
//	@Description: 获取用户信息
//	@receiver s UserServiceImpl
//	@param id 用户id
//	@return *pkg.Response 返回结果
func (s UserServiceImpl) GetUserInfo(context *gin.Context) {
	userInfo, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(-1, "无法获取用户id", err))
	}
	user := repo.NewUserRepository().SelectByID(userInfo.ID)
	if user == nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(-1, "用户不存在"))
		return
	} else {
		context.JSON(http.StatusOK, pkg.SuccessResponse(map[string]any{
			"user": user,
		}))
		return
	}
}

// SearchUsers
//
//	@Description: 查询用户
//	@receiver s UserServiceImpl
//	@param params 查询参数
//	@return *pkg.Response 返回结果
func (s UserServiceImpl) SearchUsers(context *gin.Context) {
	var params pkg.SearchUsersParams
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	params.Page = &page
	limit, _ := strconv.Atoi(context.DefaultQuery("page_size", "10"))
	params.Limit = &limit
	if context.Query("account") != "" {
		account := context.Query("account")
		params.Account = &account
	}
	if context.Query("name") != "" {
		name := context.Query("name")
		params.Name = &name
	}
	if context.Query("direction") != "" {
		direction, _ := strconv.Atoi(context.Query("direction"))
		params.Direction = &direction
	}
	users := repo.NewUserRepository().SearchUsers(params)
	pgGoalRepo := repo.NewPGGoalRepo()
	emplGoalRepo := repo.NewEmplGoalRepo()

	// 构造返回消息
	data := make([]map[string]any, 0)
	for _, user := range users {
		userInfo := make(map[string]any)
		userInfo["user"] = user
		if user.Direction == 1 {
			pgGoal, err := pgGoalRepo.SelectByUID(user.ID)
			if err != nil {
				userInfo["goal"] = "error: " + err.Error()
				continue
			}
			userInfo["goal"] = pgGoal
		} else if user.Direction == 2 {
			emplGoal, err := emplGoalRepo.SelectByUID(user.ID)
			if err != nil {
				userInfo["goal"] = "error: " + err.Error()
				continue
			}
			userInfo["goal"] = emplGoal
		} else {
			userInfo["goal"] = nil
		}
		data = append(data, userInfo)
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(data))
}
