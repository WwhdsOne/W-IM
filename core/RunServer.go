package core

import (
	"W-IM/global"
	"W-IM/initialize"
	"W-IM/middleware"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	cfg := initialize.InitConfig()
	global.Redis = initialize.InitRedisClient(cfg)
	// gin启动，初始化中间件
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.ZapLoggerMiddleware())

	initialize.InitRouter(r)
	r.Run(":8080")
}
