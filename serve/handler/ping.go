/**
 * @Author: mariomang
 * @Date: 2023-12-05 15:00:31
 * @File: handler/ping.go
**/

package handler

import "github.com/gin-gonic/gin"

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
