package pkg

import (
	"WorkProgressRecord/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtKey = []byte(config.GetServerSecretKey()) // 用于签名的密钥

// MyClaims 自定义载荷内容
type MyClaims struct {
	jwt.StandardClaims
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Permission int    `json:"permission"`
	SessionID  string `json:"session_id"`
}

// GenerateToken
//
//	@Description: 生成一个token
//	@param id 用户id
//	@param permission 用户权限
//	@param createdAt 用户创建时间
//	@param sessionID 用户sessionID
//	@return string token
//	@return error 错误
func GenerateToken(id int64, permission int, name, sessionID string) (string, error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	claims := &MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
		ID:         id,
		Permission: permission,
		Name:       name,
		SessionID:  sessionID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
