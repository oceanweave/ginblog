package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	// 大于等于2 之前默认0是管理员 1 是阅读者 但在 validate中 默认0是空值
	// 所以这里 gte=2 就是大于等于2，要所有都+1 就是 1是管理员  2是阅读者
	Role int `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// -------- 查询用户是否存在 ---------
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED // 1001
	}
	return errmsg.SUCCES
}

// --------- 新增用户 ----------
func CreateUser(data *User) int {
	// 此处密码加密 改为了下面的钩子函数 BeforeSave
	//data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

// ----------- 查询用户列表 -----------
// 为了防止获取过多，拖慢，所以分页获取
// 增加获取总数total，是为了之后更好分页，以及前端展示
func GetUsers(pageSize int, pageNum int) ([]User, int) {
	var users []User
	var total int
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// ------- 密码加密 ------------
// 此处调用钩子函数，在存入数据库之前做 密码加密
// http://v1.gorm.io/zh_CN/docs/hooks.html
// 不需要我们调用，存储前会自动调用
func (u *User) BeforeSave() {
	u.Password = ScryptPw(u.Password)
}

// 密码加密功能
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 34, 45, 68, 42, 35, 123}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// --------- 删除用户 ---------
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

// --------- 编辑用户 --------
// 此部分只能编辑基本信息，更改密码需要独立功能
// http://v1.gorm.io/zh_CN/docs/update.html
// 结构体更新形式，0值无法更新（也就是role字段不能更新），因此采用 map 方式更新
func EditUser(id int, data *User) int {
	var user User // 相当于 var user = User{}  简化了
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

// 登陆验证
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCES
}
