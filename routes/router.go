package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	// 日志中间件  记录信息等
	r.Use(middleware.Logger())
	// 跨域
	r.Use(middleware.Cors())
	// 猜测适用于 err  recovery  避免日志的错误影响后面
	r.Use(gin.Recovery())
	// 需要 bear token 的路由组
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的接口

		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 文章模块的接口
		auth.POST("art/add", v1.AddArticle)
		auth.DELETE("art/:id", v1.DeleteArt)
		// 分类模块的接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		// 上传文件
		auth.POST("upload", v1.UpLoad)
	}

	// 不需要使用 token 的路由组
	routerV1 := r.Group("api/v1")
	// 下面的第一个大括号没什么作用，就是为了美观
	{
		// 添加用户 不需要token  因为用户登录后才能获取 token
		routerV1.POST("user/add", v1.AddUser)

		routerV1.GET("users", v1.GetUsers)

		routerV1.GET("arts", v1.GetArts)
		routerV1.GET("arts/cate/:id", v1.GetCateArts)
		routerV1.GET("arts/art/:id", v1.GetArtInfo)
		routerV1.PUT("art/:id", v1.EditArt)

		routerV1.GET("categorys", v1.GetCates)
		// 登录接口  会获取 token ，需要设置 bear token  用于 auth 组的登录
		routerV1.POST("login", v1.Login)
		//routerV1.POST("upload", v1.UpLoad)
	}

	r.Run(utils.HttpPort)
}
