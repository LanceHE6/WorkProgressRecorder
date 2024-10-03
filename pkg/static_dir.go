package pkg

import (
	"WorkProgressRecord/config"
	"log"
	"os"
)

// 初始化静态文件夹
func init() {
	// 检查文件夹是否存在
	if _, err := os.Stat(config.GetExportPath()); os.IsNotExist(err) {
		// 文件夹不存在，创建文件夹
		err := os.MkdirAll(config.GetExportPath(), os.ModePerm)
		if err != nil {
			log.Fatalf("创建文件夹失败：%v", err)
		}
	}
}

// GetFullExportPath
//
//	@Description: 获取静态文件夹的完整路径
//	@param fileName 文件名
//	@return string 完整路径
func GetFullExportPath(fileName string) string {
	return config.GetExportPath() + "/" + fileName
}
