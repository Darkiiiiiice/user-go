/**
 * @Author: mariomang
 * @Date: 2024-01-18 23:49:20
 * @File: serve/serve.go
**/

package serve

import (
	"github.com/gin-gonic/gin"
	"user-go/serve/handler"
)

func RegisterRoter() *gin.Engine {

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", handler.Ping)

	}
	user := v1.Group("/user")
	{
		user.POST("", handler.CreateUser)
		user.GET("/:username", handler.GetUser)
		user.GET("", handler.ListUser)
		user.DELETE("/:username", handler.DeleteUser)
	}

	return r
}

func StartServe(r *gin.Engine, addr string) error {
	return r.Run(addr)
}
