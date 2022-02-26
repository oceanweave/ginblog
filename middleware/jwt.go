package middleware

import (
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// jwt 包 https://pkg.go.dev/github.com/dgrijalva/jwt-go/v4#StandardClaims
var JwtKey = []byte(utils.JwtKey)
var code int

type MyClaims struct {
	Username string `json:"username"`
	//Password string `json:"password"`
	jwt.StandardClaims
}

// 生成 token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaimes := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(expireTime),
			Issuer:    "ginblog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaimes)
	token, err := reqClaim.SignedString(JwtKey) // 用 JwtKey 加盐 转换为 string
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCES
}

// 验证 token
// https://pkg.go.dev/github.com/dgrijalva/jwt-go#example-Parse--Hmac
func CheckToken(token string) (*MyClaims, int) {
	// 解析 token
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		// 使用公钥去验证
		return JwtKey, nil
	})
	// 判断是否成功验证
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errmsg.SUCCES
	} else {
		return nil, errmsg.ERROR
	}
}

// jwt 中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 默认写法
		tokenHeader := c.Request.Header.Get("Authorization")
		code = errmsg.SUCCES
		// token 不存在
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			// 阻止调用后续的处理函数
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		// token 格式不对
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		// token 错误
		key, tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		// token 已过期
		if time.Now().Unix() > key.ExpiresAt.Unix() {
			code = errmsg.ERROR_TOKEN_RUNTIME

			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			// 阻止调用后续的处理函数
			c.Abort()
			return
		}
		// 可以在请求上下文里面设置一些值，然后其他地方取值
		// 可以跨中间件取值
		c.Set("username", key.Username)
		//调用后续的处理函数  因为会有很多中间件  这表示继续调用下一个中间件
		//c.Next()
	}
}
