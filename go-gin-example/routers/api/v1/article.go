package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/youminghang/go-gin-example/forms"
	"github.com/youminghang/go-gin-example/models"
	"github.com/youminghang/go-gin-example/pkg/e"
	"github.com/youminghang/go-gin-example/pkg/setting"
	"github.com/youminghang/go-gin-example/pkg/util"
)

// 获取单个文章
func GetArticle(c *gin.Context) {
	var article models.Article
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if result := setting.DB.Preload("Tag").First(&article, id); result.RowsAffected == 0 {
			code = e.ERROR_NOT_EXIST_ARTICLE
			return
		}
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": article,
	})

}

// 获取多个文章
func GetArticles(c *gin.Context) {
	valid := validation.Validation{}

	var articles []models.Article
	data := make(map[string]interface{})

	// 可以通过tag或者state查询
	maps := make(map[string]interface{})
	if arg := c.Query("tag_id"); arg != "" {
		tagId := com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}
	if arg := c.Query("state"); arg != "" {
		state := com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态值必须为0或1")
	}
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		result := setting.DB.Preload("Tag").Where(maps).Scopes(util.Paginate(util.GetPage(c), setting.ServerConfig.PageSize)).Find(&articles)
		var cnt int64
		result.Count(&cnt)
		data["total"] = cnt
		data["list"] = articles
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// 新增文章
func AddArticle(c *gin.Context) {
	articleForm := forms.AddArticleForm{}
	if err := c.ShouldBind(&articleForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	code := e.SUCCESS
	// 需要查询Tag存不存在
	if exit := ExistTagByID(articleForm.TagId); !exit {
		code = e.ERROR_NOT_EXIST_TAG
	} else {
		article := models.Article{
			TagId:     articleForm.TagId,
			Title:     articleForm.Title,
			Desc:      articleForm.Desc,
			Content:   articleForm.Content,
			CreatedBy: articleForm.CreatedBy,
			State:     *articleForm.State,
		}
		setting.DB.Save(&article)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// 修改文章
func EditArticle(c *gin.Context) {
	articleForm := forms.EditeArticleForm{}
	if err := c.ShouldBind(&articleForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	/*
		type EditeArticle struct {
			TagId     int32  `form:"tag_id" json:"tag_id" binding:"omitempty,min=1"`
			Title     string `form:"title" json:"title" binding:"omitempty,min=1,max=100"`
			Desc      string `form:"desc" json:"desc" binding:"omitempty,min=1,max=255"`
			Content   string `form:"content" json:"content" binding:"omitempty,min=1"`
			UpdatedBy string `form:"updated_by" json:"updated_by" binding:"required,min=1,max=100"`
			State     *int   `form:"state" json:"state" binding:"omitempty,oneof=0 1"`
		}
		待会记得调试一下如果没有传进来title之类的，这个结果体对应的数据会是什么
	*/
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.ERROR_NOT_EXIST_ARTICLE
	if ExistArticleByID(int32(id)) {
		maps := make(map[string]interface{})
		if articleForm.TagId != nil {
			if !ExistTagByID(*articleForm.TagId) {
				code = e.ERROR_NOT_EXIST_TAG
				return
			}
			maps["tag_id"] = articleForm.TagId
		}
		if articleForm.Title != "" {
			maps["title"] = articleForm.Title
		}
		if articleForm.Content != "" {
			maps["content"] = articleForm.Content
		}
		if articleForm.Desc != "" {
			maps["desc"] = articleForm.Desc
		}
		maps["updated_by"] = articleForm.UpdatedBy
		if articleForm.State != nil {
			maps["state"] = articleForm.State
		}
		setting.DB.Table("article").Where("id=?", id).Updates(maps)
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}

// 删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.INVALID_PARAMS
	if ExistArticleByID(int32(id)) {
		code = e.SUCCESS
		setting.DB.Where("id = ?", id).Delete(&models.Article{})
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
