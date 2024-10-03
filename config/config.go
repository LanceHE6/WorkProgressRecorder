package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

// @Description: viper配置对象
var v *viper.Viper

// init
//
//	@Description: 初始化配置
func init() {
	v = viper.New()
	v.SetConfigName("config") // 配置文件的名称（不需要带后缀）
	v.SetConfigType("yaml")   // 配置文件的类型
	v.AddConfigPath(".")      // 查找配置文件所在的路径

	// 尝试加载配置文件
	if err := v.ReadInConfig(); err != nil {
		// 如果配置文件不存在，警告用户并使用默认值
		fmt.Println("Using default config settings...")

		v.SetDefault("server.port", "8080")                                   // 服务器端口
		v.SetDefault("server.mode", "debug")                                  // 服务器日志模式
		v.SetDefault("server.log.path", "logs")                               // 服务器日志路径
		v.SetDefault("server.log.max_files", 15)                              // 服务器日志文件最大数量
		v.SetDefault("server.secretKey", "work_progress_recorder_secret_key") // 服务器密钥

		v.SetDefault("db.mysql.host", "localhost")    // mysql数据库地址
		v.SetDefault("db.mysql.port", "3306")         // mysql数据库端口
		v.SetDefault("db.mysql.account", "root")      // mysql数据库账号
		v.SetDefault("db.mysql.password", "password") // mysql数据库密码
		v.SetDefault("db.mysql.dbname", "wpr")        // mysql数据库名称

		v.SetDefault("pgee.pgee_time", "2024-12-21") // pgee时间
		v.SetDefault("export.path", "/export")       // 导出文件路径

		// 选择在此处写入默认配置文件或者不写
		err := v.WriteConfigAs("config.yaml")
		if err != nil {
			log.Fatalf("error writing default config file: %v", err)
		}
	}
}

// 获取配置,如果环境变量存在则使用环境变量，否则使用配置文件

func GetServerPort() string {
	envValue := os.Getenv("SERVER_PORT")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("server.port")
	}
}
func GetServerMode() string {
	envValue := os.Getenv("SERVER_MODE")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("server.mode")
	}
}
func GetServerLogPath() string {
	envValue := os.Getenv("SERVER_LOG_PATH")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("server.log.path")
	}
}
func GetServerLogMaxFiles() int {
	envValue := os.Getenv("SERVER_LOG_MAX_FILES")
	if envValue != "" {
		return v.GetInt("server.log.max_files")
	} else {
		return v.GetInt("server.log.max_files")
	}
}
func GetServerSecretKey() string {
	envValue := os.Getenv("SERVER_SECRET_KEY")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("server.secretKey")
	}
}
func GetDBMySQLPort() string {
	envValue := os.Getenv("MYSQL_PORT")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("db.mysql.port")
	}
}
func GetDBMySQLAccount() string {
	envValue := os.Getenv("MYSQL_ACCOUNT")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("db.mysql.account")
	}
}
func GetDBMySQLPassword() string {
	envValue := os.Getenv("MYSQL_PASSWORD")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("db.mysql.password")
	}
}
func GetDBMySQLDBName() string {
	envValue := os.Getenv("MYSQL_DBNAME")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("db.mysql.dbname")
	}
}
func GetPgeeTime() string {
	envValue := os.Getenv("PGEETIME")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("pgee.pgee_time")
	}
}
func GetDBMySQLHost() string {
	envValue := os.Getenv("MYSQL_HOST")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("db.mysql.host")
	}
}
func GetExportPath() string {
	envValue := os.Getenv("EXPORT_PATH")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("export.path")
	}
}
