package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	// 设置配置文件名称（不包括文件扩展名）
	viper.SetConfigName("config")

	// 设置配置文件类型
	viper.SetConfigType("toml")

	// 设置配置文件搜索路径，默认为项目的根目录，可以添加多个路径进行搜索
	// viper.AddConfigPath("$GOPATH/src/github.com/username/project/config")
	// viper.AddConfigPath("/etc/appname/")          // 添加系统配置路径
	// viper.AddConfigPath("/usr/local/etc/appname/") // 添加本地配置路径
	// viper.AddConfigPath(".")                     // 添加当前工作目录
	viper.AddConfigPath(".")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 获取配置项
	message := viper.GetString("message")
	fmt.Println("获取配置文件中Message的值=", message)

}
