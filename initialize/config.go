package initialize

import (
	"W-IM/config"
	"github.com/spf13/viper"
	"log"
)

func InitConfig() *config.Config {
	// 初始化 Viper
	viper.SetConfigName("conf/app")
	viper.AddConfigPath(".")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 创建配置实例
	var cfg config.Config

	// 将配置映射到结构体
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
	return &cfg
}
