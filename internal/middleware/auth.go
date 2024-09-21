package middleware

import (
	"WorkProgressRecord/internal/repo"
	"WorkProgressRecord/pkg"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware 基础鉴权的中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, pkg.NewResponse(11, "未提供 Authorization 鉴权头", nil))
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, pkg.NewResponse(12, "非法token", nil))
			c.Abort()
			return
		}
		myClaims, err := GetUserInfoByContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, pkg.NewResponse(13, "非法token", err.Error()))
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(bearerToken[1], &pkg.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return pkg.JwtKey, nil
		})

		if err != nil {
			var ve *jwt.ValidationError
			if errors.As(err, &ve) {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					// Token is expired
					c.JSON(http.StatusUnauthorized, pkg.NewResponse(14, "token已过期", nil))
				} else {
					// Other errors
					c.JSON(http.StatusUnauthorized, pkg.NewResponse(15, "非法token", nil))
				}
			}
			c.Abort()
			return
		}

		// 判断是否在数据库中
		userRepo := repo.NewUserRepository()
		user := userRepo.SelectByID(myClaims.ID)
		if user == nil {
			c.JSON(http.StatusUnauthorized, pkg.NewResponse(16, "用户不存在", nil))
			c.Abort()
		}
		if user.SessionID != myClaims.SessionID {
			c.JSON(http.StatusUnauthorized, pkg.NewResponse(17, "token已失效", nil))
			c.Abort()
		}

		if _, ok := token.Claims.(*pkg.MyClaims); ok && token.Valid {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, pkg.NewResponse(18, "非法token", nil))
			c.Abort()
			return
		}
	}
}

// IsAdminMiddleware 管理员鉴权中间件 >=2
//
//	@Description: 管理员鉴权中间件
//	@return gin.HandlerFunc
func IsAdminMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		AuthMiddleware()
		// 根据token判断permission是否大于等于2
		myClaims, err := GetUserInfoByContext(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, pkg.NewResponse(19, "非法token", nil))
			context.Abort()
			return
		}

		if myClaims.Permission >= 2 {
			context.Next()
		} else {
			context.JSON(http.StatusForbidden, pkg.NewResponse(1, "权限拒绝", nil))
			context.Abort()
			return
		}
	}
}

// GetUserInfoByContext
//
//	@Description: 从context中获取用户信息
//	@param context *gin.Context
//	@return pkg.MyClaims 用户信息
//	@return error 错误信息
func GetUserInfoByContext(context *gin.Context) (pkg.MyClaims, error) {
	authHeader := context.GetHeader("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	// 解析token
	claims := pkg.MyClaims{}
	_, err := jwt.ParseWithClaims(bearerToken[1], &claims, func(token *jwt.Token) (interface{}, error) {
		return pkg.JwtKey, nil
	})
	// 从token中获取载荷数据
	return claims, err
}
