package routes

import "github.com/gin-gonic/gin"

func RegisterApiRoute(route *gin.Engine) {
	api := route.Group("/api")
	{
		api.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"ping": "pong",
			})
		})
	}
}
