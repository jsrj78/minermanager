/*
 * @Author: chunhua yang
 * @Date: 2023-10-23 22:39:46
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-23 22:49:52
 * @FilePath: /minermanager/server/controllers/base.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (ctrl *BaseController) Response(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func (ctl *BaseController) ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   data,
	})
}

func (ctrl *BaseController) ResponseError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status": "error",
		"data":   message,
	})
}
