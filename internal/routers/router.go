package routers

import (
	v1 "blog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := v1.NewArticle()
	tag := v1.NewTag()

	apiv1 := r.Group("api/v1")
	{
		apiv1.POST("tags", tag.Create)
		apiv1.GET("tags", tag.List)
		apiv1.GET("tags/:id", tag.Get)
		apiv1.PUT("tags/:id", tag.Update)

		apiv1.POST("articles", article.Create)
		apiv1.GET("articles", article.List)
		apiv1.GET("articles/:id", article.Get)
		apiv1.PUT("articles/:id", article.Update)
	}

	return r
}
