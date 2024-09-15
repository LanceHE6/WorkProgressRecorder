package pkg

import "github.com/google/uuid"

// GenerateUUID
//
//	@Description: 生成指定长度的UUID
//	@param length 长度,为-1时，生成默认长度32的UUID
//	@return string UUID
func GenerateUUID(length int) string {
	// length为-1时，生成默认长度32的UUID
	newUUID := uuid.New().String()
	if length == -1 {
		return newUUID
	}
	return newUUID[len(newUUID)-length:]
}
