package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"ginblog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// 查询用户是否存在
func UserExist(c *gin.Context) {

}

// 添加用户
func AddUser(c *gin.Context) {
	// 先查询用户是否存在
	var data model.User
	var msg string
	_ = c.ShouldBindJSON(&data)
	// 添加用户前进行 信息验证
	// 通过设置 model User 的 validate tag 进行格式要求，最小不能少于 最大不能多于多少个字符等
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCES {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	code = model.CheckUser(data.Username)
	// 创建用户
	if code == errmsg.SUCCES {
		model.CreateUser(&data)
	}
	//if code == errmsg.ERROR_USERNAME_USED {
	//	code = errmsg.ERROR_USERNAME_USED
	//}
	// 返回相应 穿件成功
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		// gorm 的 -1 表示不限制
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCES
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
// todo: 应该先查询用户是否存在，目前bug是：删除不存在的用户，也会返回 ok
// 软删除：在Navicat数据表中仍会看到数据，但是查询用户是查询不到的。也就是删除并不会马上删除真实数据，但会标记删除，使其无法使用
func DeleteUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
// todo：问题：只有更新用户名才能更新，因为是查询用户名是否存在
// todo: 只更新用户信息，不更新用户名情况
func EditUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	// 用户名存在显示 1001 错误
	code = model.CheckUser(data.Username)
	// 更新用户
	if code == errmsg.SUCCES {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
