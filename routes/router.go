package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	// 下面的第一个大括号没什么作用，就是为了美观
	{
		// 用户模块的接口
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		routerV1.PUT("user/:id", v1.EditUser)
		routerV1.DELETE("user/:id", v1.DeleteUser)
		// 文章模块的接口
		routerV1.POST("art/add", v1.AddArticle)
		routerV1.GET("arts", v1.GetArts)
		routerV1.GET("arts/cate/:id", v1.GetCateArts)
		routerV1.GET("arts/art/:id", v1.GetArtInfo)
		routerV1.PUT("art/:id", v1.EditArt)
		routerV1.DELETE("art/:id", v1.DeleteArt)
		// 分类模块的接口
		routerV1.POST("category/add", v1.AddCategory)
		routerV1.GET("categorys", v1.GetCates)
		routerV1.PUT("category/:id", v1.EditCate)
		routerV1.DELETE("category/:id", v1.DeleteCate)
	}

	r.Run(utils.HttpPort)
}
