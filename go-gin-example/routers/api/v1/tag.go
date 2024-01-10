package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youminghang/go-gin-example/models"
	"github.com/youminghang/go-gin-example/pkg/e"
	"github.com/youminghang/go-gin-example/pkg/setting"
)

// 获取多个文章标签
func GetTags(c *gin.Context) {
	var tags []models.Tag
	data := make(map[string]interface{})
	code := e.SUCCESS
	name := c.Query("name")
	result := setting.DB.Model(&models.Tag{}).Where("name LIKE ?", "%"+name+"%").Find(&tags)
	if result.RowsAffected == 0 {
		code = e.ERROR_EXIST_TAG
		c.JSON(http.StatusNotFound, gin.H{
			"msg": e.GetMsg(code),
		})
	}
	var cnt int64
	result.Count(&cnt)
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
}

// 修改文章标签
func EditTag(c *gin.Context) {

}

// 删除文章标签
func DeleteTag(c *gin.Context) {
}
