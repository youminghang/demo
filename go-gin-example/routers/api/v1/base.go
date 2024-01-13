package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/youminghang/go-gin-example/models"
	"github.com/youminghang/go-gin-example/pkg/e"
	"github.com/youminghang/go-gin-example/pkg/setting"
)

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for filed, err := range fileds {
		rsp[filed[strings.Index(filed, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorError(c *gin.Context, err error) {
	code := e.INVALID_PARAMS
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code) + errs.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"msg":  removeTopStruct(errs.Translate(setting.Trans)),
	})
	return
}

func ExistTagByID(id int32) bool {
	var tag models.Tag
	setting.DB.Select("id").Where("id = ?", id).First(&tag)

	if tag.ID > 0 {
		return true
	}

	return false
}

func ExistArticleByID(id int32) bool {
	var article models.Article
	setting.DB.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}
