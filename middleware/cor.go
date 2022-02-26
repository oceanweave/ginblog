package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			//AllowAllOrigins: true, // 与下面二选一
			AllowOrigins:  []string{"*"}, // 允许哪个域名过来 跨域
			AllowMethods:  []string{"*"}, // 允许所有方法
			AllowHeaders:  []string{"Origin"},
			ExposeHeaders: []string{"Content-Length", "Authorzation"},
			//AllowCredentials: true, // 是不是发送 cookie 请求
			//AllowOriginFunc: func(origin string) bool {
			//	return origin == "https://github.com"
			//},
			MaxAge: 12 * time.Hour, // 域请求持续时间
		})
	}
}
