package ginapi

import (
	"github.com/gin-gonic/gin"
)

func NewGinServer() error {
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.GET("/gin/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "get,hello gin",
			})
		})
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/gin/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "post,hello gin",
			})
			panic("test")
		})
	}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	return r.Run(":4444")
}
