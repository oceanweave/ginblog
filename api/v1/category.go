package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 查询分类是否存在

// 添加分类

// 查询单个分类下的文章

// 查询分类列表

// 编辑分类

// 删除分类

// 查询分类是否存在
func CategoryExist(c *gin.Context) {

}

// 添加分类
func AddCategory(c *gin.Context) {
	// 先查询分类是否存在
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	// 创建分类
	if code == errmsg.SUCCES {
		model.CreateCate(&data)
	}
	//if code == errmsg.ERROR_CategoryNAME_USED {
	//	code = errmsg.ERROR_CategoryNAME_USED
	//}
	// 返回相应 穿件成功
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个分类

// todo: 查询分类下的所有文章

// 查询分类列表
func GetCates(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		// gorm 的 -1 表示不限制
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetCates(pageSize, pageNum)
	code = errmsg.SUCCES
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除分类
// todo: 应该先查询分类是否存在，目前bug是：删除不存在的分类，也会返回 ok
// 软删除：在Navicat数据表中仍会看到数据，但是查询分类是查询不到的。也就是删除并不会马上删除真实数据，但会标记删除，使其无法使用
func DeleteCate(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑分类
// todo：问题：只有更新分类名才能更新，因为是查询分类名是否存在
func EditCate(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	// 分类名存在显示 2001 错误
	code = model.CheckCategory(data.Name)
	// 更新分类
	if code == errmsg.SUCCES {
		model.EditCate(id, &data)
	}
	if code == errmsg.ERROR_CATEGORY_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
