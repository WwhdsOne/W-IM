package initialize

import (
	"W-IM/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	group := r.Group("/api")
	api.InitTalkRouter(group)
}
