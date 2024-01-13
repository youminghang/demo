package api

import (
	"log"
	"net/http"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/youminghang/go-gin-example/middlewares"
	"github.com/youminghang/go-gin-example/pkg/e"
	"github.com/youminghang/go-gin-example/pkg/util"
	v1 "github.com/youminghang/go-gin-example/routers/api/v1"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// 用于生成token
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		if !v1.CheckAuth(username, password) {
			code = e.ERROR_AUTH
			return
		}
		j := middlewares.NewJWT()
		claims := util.CustomClaims{
			Username: a.Username,
			Password: a.Password,
			StandardClaims: jwt.StandardClaims{
				NotBefore: time.Now().Unix(),              // 签名的生效时间
				ExpiresAt: time.Now().Unix() + 60*60*24*7, // 设置7天过期
				Issuer:    "imooc",
			},
		}
		token, err := j.CreateToken(claims)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		}
		data["token"] = token
		code = e.SUCCESS

	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
