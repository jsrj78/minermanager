/*
 * @Author: chunhua yang
 * @Date: 2023-10-22 18:56:28
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-23 22:56:01
 * @FilePath: /minermanager/server/routes/api.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */
package routes

import (
	"mmserver/controllers"

	"github.com/gin-gonic/gin"
)

var userCtrl controllers.UserController

func InitApiRouterGroup(r *gin.RouterGroup) {
	r.GET("/user", userCtrl.GetAllUserByPage)
}
