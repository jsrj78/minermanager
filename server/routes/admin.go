/*
 * @Author: chunhua yang
 * @Date: 2023-10-22 18:56:33
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-23 22:54:49
 * @FilePath: /minermanager/server/routes/admin.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */
package routes

import (
	"mmserver/controllers"

	"github.com/gin-gonic/gin"
)

var SysController controllers.SysController

func InitAdminRouterGroup(r *gin.RouterGroup) {

	r.POST("/shelf/autoadd", SysController.AddShelfByRule)
}
