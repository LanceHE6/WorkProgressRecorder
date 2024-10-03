package impl

import (
	"WorkProgressRecord/internal/middleware"
	"WorkProgressRecord/internal/model"
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"net/http"
	"strconv"
	"time"
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
		Class   string `json:"class" form:"class" binding:"required"`
		Major   string `json:"major" form:"major" binding:"required"`
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
		user.Class = userInfo.Class
		user.Major = userInfo.Major
		user.Direction = model.Direction{}
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
		ID          int64  `json:"id" form:"id"`
		OldPassword string `json:"old_password" form:"old_password"`
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
	// 修改他人密码
	if data.ID != 0 && data.ID != claims.ID {
		// 非管理员权限
		if claims.Permission != 2 {
			context.JSON(http.StatusBadRequest, pkg.FailedResponse(1, "权限拒绝"))
			return
		}
		err = repo.NewUserRepository().UpdatePassword(data.ID, data.NewPassword)
		if err != nil {
			context.JSON(http.StatusInternalServerError, pkg.ErrorResponse(-3, "修改密码失败", err))
			return
		} else {
			context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
			return
		}
	} else {
		// 修改自己密码
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
	if user.Direction == nil {
		context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
		return
	}
	goalData := map[string]any{
		"pg_goal":   nil,
		"empl_goal": nil,
	}

	// 考研为目标
	if pkg.IsContainGoal(user.Direction, model.Postgraduate) {
		pgGoal, err := repo.NewPGGoalRepo().SelectByUID(userInfo.ID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, pkg.ErrorResponse(-2, "查询考研目标失败", err))
			return
		}
		goalData["pg_goal"] = pgGoal
	}
	// 就业为目标
	if pkg.IsContainGoal(user.Direction, model.Employment) {
		emplGoal, err := repo.NewEmplGoalRepo().SelectByUID(userInfo.ID)
		if err != nil {
			context.JSON(http.StatusInternalServerError, pkg.ErrorResponse(-2, "查询就业目标失败", err))
			return
		}
		goalData["empl_goal"] = emplGoal
	}
	resultData := map[string]any{
		"direction": user.Direction,
		"goal":      goalData,
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(resultData))
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
	if context.Query("keyword") != "" {
		keyword := context.Query("keyword")
		params.Keyword = &keyword
	}

	users, count := repo.NewUserRepository().SearchUsers(params)
	pgGoalRepo := repo.NewPGGoalRepo()
	emplGoalRepo := repo.NewEmplGoalRepo()

	// 构造返回消息
	data := make([]map[string]any, 0)

	for _, user := range users {
		userInfo := make(map[string]any)
		userInfo["user"] = user
		goalData := map[string]any{
			"pg_goal":   nil,
			"empl_goal": nil,
		}
		if pkg.IsContainGoal(user.Direction, model.Postgraduate) {
			pgGoal, err := pgGoalRepo.SelectByUID(user.ID)
			if err != nil {
				userInfo["goal"] = "error: " + err.Error()
			} else {
				goalData["pg_goal"] = pgGoal
			}
		}
		if pkg.IsContainGoal(user.Direction, model.Employment) {
			emplGoal, err := emplGoalRepo.SelectByUID(user.ID)
			if err != nil {
				userInfo["goal"] = "error: " + err.Error()

			} else {
				goalData["empl_goal"] = emplGoal
			}
		}
		userInfo["goal"] = goalData
		data = append(data, userInfo)
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(map[string]any{
		"rows":      data,
		"total":     count,
		"page":      page,
		"page_size": limit,
	}))
}

// UpdateUserInfo
//
//	@Description: 更新用户信息
//	@receiver s
//	@param context
func (s UserServiceImpl) UpdateUserInfo(context *gin.Context) {
	type updateUserInfo struct {
		Class string `json:"class" form:"class" binding:"required"`
		Major string `json:"major" form:"major" binding:"required"`
	}
	var data updateUserInfo
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}
	claims, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(-1, err.Error()))
	}
	userRepo := repo.NewUserRepository()
	user := userRepo.SelectByID(claims.ID)
	if user == nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(-1, "用户不存在"))
		return
	}

	user.Class = data.Class
	user.Major = data.Major
	err = userRepo.UpdateUserInfo(claims.ID, *user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, pkg.ErrorResponse(1, "更新失败", err))
		return
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
}

// DeleteUserDirection
//
//	@Description:  删除用户方向
//	@receiver s
//	@param context
func (s UserServiceImpl) DeleteUserDirection(context *gin.Context) {
	type deleteUserDirectionReq struct {
		Direction int `json:"direction" form:"direction" binding:"required"`
	}
	var data deleteUserDirectionReq
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(100, err.Error()))
		return
	}
	claims, err := middleware.GetUserInfoByContext(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(-1, err.Error()))
	}
	userRepo := repo.NewUserRepository()
	user := userRepo.SelectByID(claims.ID)
	if user == nil {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(-1, "用户不存在"))
		return
	}
	var direction model.DirectionType
	if data.Direction == 1 {
		direction = model.Postgraduate
	} else if data.Direction == 2 {
		direction = model.Employment
	} else {
		context.JSON(http.StatusBadRequest, pkg.FailedResponse(2, "未知的方向"))
		return
	}
	err = userRepo.DeleteDirection(user.ID, direction)
	if err != nil {
		context.JSON(http.StatusBadRequest, pkg.ErrorResponse(1, "删除失败", err))
		return
	}
	context.JSON(http.StatusOK, pkg.SuccessResponse(nil))
}

// ExportUserInfo2Excel
//
//	@Description: 将用户信息导出为excel文件
//	@receiver s UserServiceImpl
//	@param context *gin.Context
func (s UserServiceImpl) ExportUserInfo2Excel(context *gin.Context) {
	userRepo := repo.NewUserRepository()
	pgGoalRepo := repo.NewPGGoalRepo()
	emplGoalRepo := repo.NewEmplGoalRepo()
	users := userRepo.SelectAll()

	// 表数据
	data := make([][]string, 0)

	for _, user := range users {
		// 构造用户数据
		userInfo := make([]string, 13)
		userInfo[0] = user.Account
		userInfo[1] = user.Class
		userInfo[2] = user.Name
		userInfo[3] = user.Major
		// 构造用户方向信息
		isPg := pkg.IsContainGoal(user.Direction, model.Postgraduate)
		isEmpl := pkg.IsContainGoal(user.Direction, model.Employment)
		if isPg && isEmpl {
			userInfo[4] = "考研&就业"
		} else if isEmpl {
			userInfo[4] = "就业"
		} else if isPg {
			userInfo[4] = "考研"
		} else {
			userInfo[4] = "未确定"
		}
		if isPg {
			pgGoal, err := pgGoalRepo.SelectByUID(user.ID)
			if err != nil {
				userInfo[5] = err.Error()
				userInfo[6] = err.Error()
				userInfo[7] = err.Error()
			} else {
				userInfo[5] = pgGoal.TargetUniversity
				userInfo[6] = pgGoal.TargetMajor
				userInfo[7] = strconv.FormatFloat(pgGoal.TargetScore, 'E', -1, 64)
			}
		} else {
			userInfo[5] = ""
			userInfo[6] = ""
			userInfo[7] = ""
		}
		if isEmpl {
			emplGoal, err := emplGoalRepo.SelectByUID(user.ID)
			if err != nil {
				userInfo[8] = err.Error()
				userInfo[9] = err.Error()
				userInfo[10] = err.Error()
				userInfo[11] = err.Error()
				userInfo[12] = err.Error()
			} else {
				if emplGoal.Status == 1 {
					userInfo[8] = "未拿到offer"
				} else {
					userInfo[8] = "已拿到offer"
				}
				userInfo[9] = emplGoal.TargetCompany
				userInfo[10] = emplGoal.TargetJob
				userInfo[11] = emplGoal.TargetSalary
				userInfo[12] = emplGoal.TargetArea
			}
		} else {
			userInfo[8] = ""
			userInfo[9] = ""
			userInfo[10] = ""
			userInfo[11] = ""
			userInfo[12] = ""
		}
		data = append(data, userInfo)
	}
	// 创建新的excel文件
	excel := excelize.NewFile()
	// 设置列宽
	_ = excel.SetColWidth("Sheet1", "A", "M", 20)
	// 设置表头
	titleList := []string{"学号", "班级", "姓名", "专业", "方向", "目标院校", "目标专业", "目标分数", "找工作状态", "目标公司", "目标职位", "理想薪资", "目标地区"}
	_ = excel.SetSheetRow("Sheet1", "A1", &titleList)
	// 循环遍历写入数据
	for i, row := range data {
		rowNum := fmt.Sprintf("A%d", i+2)
		_ = excel.SetSheetRow("Sheet1", rowNum, &row)
	}
	// 命名
	fileName := "eui" + strconv.FormatInt(time.Now().Unix(), 10) + ".xlsx"
	path := pkg.GetFullExportPath(fileName)
	err := excel.SaveAs(path)
	if err != nil {
		context.JSON(http.StatusInternalServerError, pkg.ErrorResponse(1, "导出文件失败", err))
		return
	}
	// 设置响应头部
	context.Header("Content-Type", "application/octet-stream")
	context.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	err = excel.Write(context.Writer)
	if err != nil {
		context.JSON(http.StatusInternalServerError, pkg.ErrorResponse(2, "导出文件失败", err))
	}
	//context.JSON(http.StatusOK, pkg.SuccessResponse(map[string]any{
	//	"fileName": fileName,
	//	"path":     path,
	//}))
}
