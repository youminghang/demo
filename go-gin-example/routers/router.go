package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youminghang/go-gin-example/pkg/setting"
	v1 "github.com/youminghang/go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerConfig.RunMode)
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	apiv1 := r.Group("api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}
