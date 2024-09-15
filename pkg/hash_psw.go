package pkg

import (
	"crypto/sha256"
	"fmt"
	"hash"
)

// HashPsw
//
//	@Description: 哈希密码
//	@param psw 密码
//	@return string 哈希值
func HashPsw(psw string) string {
	var hashInstance hash.Hash // 定义哈希实例

	hashInstance = sha256.New()

	hashInstance.Write([]byte(psw)) // 将字符串转换为字节数组，写入哈希对象

	bytes := hashInstance.Sum(nil)  // 哈希值追加到参数后面，只获取原始值，不用追加，用nil，返回哈希值字节数组
	return fmt.Sprintf("%x", bytes) // 格式化输出哈希值
}

// CheckPsw
//
//	@Description: 校验密码
//	@param hashedPsw 哈希值
//	@param psw 密码
//	@return bool 校验结果
func CheckPsw(hashedPsw, psw string) bool {
	return hashedPsw == HashPsw(psw)
}
