/*
 * @Author: chunhua yang
 * @Date: 2023-10-23 22:40:08
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-29 00:27:28
 * @FilePath: /minermanager/server/controllers/user.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */
package controllers

import (
	"fmt"
	"mmserver/models"
	"mmserver/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (ctrl *UserController) GetAllUserByPage(c *gin.Context) {
	user := []models.TUser{}

	utils.DB.Find(&user)

	msm := models.NewMinerFactory().MinerStatusManager

	msm.Mux.Lock()
	for _, value := range msm.OnlineList {
		//fmt.Println("start to read online")
		fmt.Printf("在线: %s\n", value)
	}

	for _, value := range msm.OfflineList {
		//fmt.Println("start to read offline")
		fmt.Printf("离线: %s\n", value)
	}

	msm.Mux.Unlock() // 解锁访问

	time.Sleep(time.Second * 3) // 可能需

	ctrl.ResponseSuccess(c, user)

}

func (ctrl *UserController) GetUserById(c *gin.Context) {

}

func (ctrl *UserController) CreateUser(c *gin.Context) {

}

func (ctrl *UserController) UpdateUser(c *gin.Context) {

}

func (ctrl *UserController) DeleteUser(c *gin.Context) {

}
