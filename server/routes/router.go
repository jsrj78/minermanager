/*
 * @Author: chunhua yang
 * @Date: 2023-10-22 14:32:35
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-23 22:55:01
 * @FilePath: /minermanager/server/routes/router.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */
package routes

import (
	"mmserver/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    200,
		})
	})

	api := r.Group(utils.AppConfig.Api.Path + "/" + utils.AppConfig.Api.Version)
	InitApiRouterGroup(api)

	admin := r.Group("admin")
	InitAdminRouterGroup(admin)

	return r
}
