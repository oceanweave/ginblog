package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// --------- 查询分类是否存在 -----------
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED // 1001
	}
	return errmsg.SUCCES
}

// --------- 新增分类 -----------
func CreateCate(data *Category) int {
	// 此处密码加密 改为了下面的钩子函数 BeforeSave
	//data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

// --------- 查询分类列表 ----------
// 为了防止获取过多，拖慢，所以分页获取
func GetCates(pageSize int, pageNum int) []Category {
	var cates []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cates
}

// -------- todo: 查询分类下的所有文章 ---------

// --------- 删除分类 ---------
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

// --------- 编辑分类 --------
// 此部分只能编辑基本信息，更改密码需要独立功能
// http://v1.gorm.io/zh_CN/docs/update.html
// 结构体更新形式，0值无法更新（也就是role字段不能更新），因此采用 map 方式更新
func EditCate(id int, data *Category) int {
	var cate Category // 相当于 var user = User{}  简化了
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}
