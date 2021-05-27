package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	Config.SetConfigName("config")   // 配置文件名称(无扩展名)
	Config.SetConfigType("yaml")     // 如果配置文件的名称中没有扩展名，则需要配置此项
	Config.AddConfigPath("./config") // 查找配置文件所在的路径
	Config.WatchConfig()             // 启动配置热更新

	err := Config.ReadInConfig() // 查找并读取配置文件
	if err != nil {              // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// debug
	fmt.Println(Config.GetString("database.mysql.host"))
}
