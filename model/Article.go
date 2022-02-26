package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	// 指定外键为 cid ，外键就是用来关联的，就是 Category 主键 id 和 外键 cid 对应关联
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	// cid 是文章id
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// --------- 新增文章 -----------
func CreateArt(data *Article) int {
	// 此处密码加密 改为了下面的钩子函数 BeforeSave
	//data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

// todo 查询分类下的所有文章
// 增加获取文章总数total，是为了之后更好分页，以及前端展示
func GetCateArts(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var articleList []Article
	// todo 查询不存在的分类，没有返回错误
	// todo 先验证分类是否存在，不存在就报错
	var total int
	err = db.Preload("Category").Where("cid = ?", id).Find(&articleList).Limit(pageSize).Offset((pageNum - 1) * pageSize).Count(&total).Error
	if err != nil {
		return articleList, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return articleList, errmsg.SUCCES, total
}

// todo 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err = db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCES
}

// --------- 查询文章列表 ----------
// 为了防止获取过多，拖慢，所以分页获取
// 增加获取文章总数total，是为了之后更好分页，以及前端展示
func GetArts(pageSize int, pageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int
	// Preload 方法的参数应该是主结构体的字段名  主结构体 就是 父结构体
	// http://v1.gorm.io/zh_CN/docs/preload.html
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCES, total
}

// --------- 删除文章 ---------
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

// --------- 编辑文章 --------
// 此部分只能编辑基本信息，更改密码需要独立功能
// http://v1.gorm.io/zh_CN/docs/update.html
// 结构体更新形式，0值无法更新（也就是role字段不能更新），因此采用 map 方式更新
func EditArt(id int, data *Article) int {
	var art Article // 相当于 var user = User{}  简化了
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&art).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}
