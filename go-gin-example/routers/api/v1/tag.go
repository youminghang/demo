package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/unknwon/com"
	"github.com/youminghang/go-gin-example/forms"
	"github.com/youminghang/go-gin-example/models"
	"github.com/youminghang/go-gin-example/pkg/e"
	"github.com/youminghang/go-gin-example/pkg/setting"
)

// 获取多个文章标签
func GetTags(c *gin.Context) {
	var tags []models.Tag
	maps := make(map[string]interface{})
	var cnt int64
	data := make(map[string]interface{})

	code := e.SUCCESS
	name := c.Query("name")
	var pageNum int = 0
	if arg := c.Query("page"); arg != "" {
		pageNum = com.StrTo(arg).MustInt()
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	maps["name"] = name
	tags, cnt = models.GetTags(pageNum, setting.ServerConfig.PageSize, maps)

	if cnt == 0 {
		code = e.ERROR_NOT_EXIST_TAG
		c.JSON(http.StatusNotFound, gin.H{
			"msg": e.GetMsg(code),
		})
	}

	data["total"] = cnt
	data["lists"] = tags

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增文章标签
func AddTag(c *gin.Context) {
	tagForm := forms.AddTagForm{}
	if err := c.ShouldBind(&tagForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	var tag models.Tag
	code := e.SUCCESS
	if result := setting.DB.Model(&models.Tag{}).Where("name LIKE ?", tagForm.Name).First(&tag); result.RowsAffected != 0 {
		code = e.ERROR_EXIST_TAG
	} else {
		tag = models.Tag{
			Name:      tagForm.Name,
			CreatedBy: tagForm.CreatedBy,
			State:     *tagForm.State,
		}
		setting.DB.Model(&models.Tag{}).Save(&tag)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}

// 修改文章标签
func EditTag(c *gin.Context) {
	tagForm := forms.UpdateTagForm{}
	if err := c.ShouldBind(&tagForm); err != nil {
		HandleValidatorError(c, err)
		return
	}

	id := com.StrTo(c.Param("id")).MustInt()
	var tag models.Tag
	code := e.SUCCESS
	// 查询这个id的tag是否存在
	if result := setting.DB.First(&tag, id); result.RowsAffected == 0 {
		code = e.ERROR_NOT_EXIST_TAG
	} else {
		tag.Name = tagForm.Name
		tag.UpdatedBy = tagForm.UpdatedBy
		if tagForm.State != nil {
			tag.State = *tagForm.State
		}
		setting.DB.Save(&tag)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}

// 删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	var tag models.Tag
	code := e.SUCCESS
	// 查询这个id的tag是否存在
	if result := setting.DB.First(&tag, id); result.RowsAffected == 0 {
		code = e.ERROR_NOT_EXIST_TAG
	} else {
		setting.DB.Delete(&tag,id)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

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
			"msg":  err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"msg":  removeTopStruct(errs.Translate(setting.Trans)),
	})
	return
}
